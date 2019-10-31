package resourceScheduler

import (
	"backend/util"
	"context"
	"errors"
	"github.com/satori/go.uuid"
	"sync"
)

type BasicScheduler struct {
	//typeResPool map[Type]*list.List
	uuidResPool map[uuid.UUID]Resource
	requests    map[chan []Resource][]Resource //put something to wakeup request then actual []Resource, so send twice
	mutex       sync.Mutex
}

func NewBasicScheduler() Scheduler {
	bs := BasicScheduler{}
	bs.init()
	return &bs
}

func (s *BasicScheduler) init() {
	s.uuidResPool = map[uuid.UUID]Resource{}
	s.requests = map[chan []Resource][]Resource{}
}

func (s *BasicScheduler) RegRes(resource Resource) util.Err {
	if s.uuidResPool == nil || s.requests == nil {
		return NewUseOfNoneInitScheduler(nil)
	}
	_, id := resource.Definition()
	if id == uuid.Nil {
		return NewIllegalUUID(errors.New("uuid is zero, please make sure resource have unique uuid"))
	}
	s.uuidResPool[id] = resource
	return nil
}

func (s *BasicScheduler) Request(ctx context.Context, logic func(), resources ...Resource) util.Err {
	resources, err := s.safeGetAllRes(ctx, resources...)
	if err != nil {
		return err
	}
	logic()
	s.safeReturnAllRes(resources)
	return nil
}

//return if all res is accessible in the pool
//will return NewInvalidResourceError if refereed res did not reg to this scheduler
func (s *BasicScheduler) resAllAccessible(resources ...Resource) (bool, util.Err) {
	for _, resource := range resources {
		_, id := resource.Definition()
		if id != uuid.Nil {
			if res, ok := s.uuidResPool[id]; !ok {
				return false, NewInvalidResourceError(nil)
			} else if res == nil {
				return false, nil
			}
		} else {
			return false, NewIllegalUUID(nil)
		}
	}
	return true, nil
}

func (s *BasicScheduler) getAllRes(resources ...Resource) []Resource {
	var gotRes []Resource
	for _, resource := range resources {
		_, id := resource.Definition()
		gotRes = append(gotRes, s.uuidResPool[id])
		s.uuidResPool[id] = nil
	}
	return gotRes
}

//wait until get all res or ctx canceled
func (s *BasicScheduler) safeGetAllRes(ctx context.Context, resources ...Resource) ([]Resource, util.Err) {
	s.mutex.Lock()
	if ok, err := s.resAllAccessible(resources...); err != nil {
		s.mutex.Unlock()
		return nil, err
	} else if ok {
		gotRes := s.getAllRes(resources...)
		s.mutex.Unlock()
		return gotRes, nil
	} else {
		wakeUp := make(chan []Resource)
		s.requests[wakeUp] = resources
		s.mutex.Unlock()
		select {
		case <-ctx.Done():
			close(wakeUp)
			return nil, util.NewContextCanceled(nil)
		case <-wakeUp:
			return <-wakeUp, nil
		}
	}
}

//return those occupied res and let others Requests to run
func (s *BasicScheduler) safeReturnAllRes(resources []Resource) {
	var wakeUp = func(ch chan []Resource) (success bool) {
		defer func() {
			//will panic if the request already canceled
			if p := recover(); p != nil {
				success = false
			}
			delete(s.requests, ch) //delete the request despite if request canceled
		}()
		ch <- nil
		return true
	}
	go func() {
		s.mutex.Lock()
		for _, resource := range resources {
			if resource!=nil{	//resource can be nil if request contains same res for more than one times!
				_, id := resource.Definition()
				s.uuidResPool[id] = resource
			}
		}
		for ch, reqRes := range s.requests {
			if ok, _ := s.resAllAccessible(reqRes...); ok {
				if wakeUp(ch) {
					ch <- s.getAllRes(reqRes...)
				}
			}
		}
		s.mutex.Unlock()
	}()
}
