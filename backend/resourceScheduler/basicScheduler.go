package resourceScheduler

import (
	"backend/util"
	"container/list"
	"context"
	"github.com/satori/go.uuid"
	"sync"
)

type BasicScheduler struct {
	typeResPool map[Type]*list.List
	uuidResPool map[uuid.UUID]Resource
	requests    map[chan []Resource][]Resource	//put something to wakeup request then actual []Resource, so send twice
	mutex       sync.Mutex
}

func NewBasicScheduler() Scheduler {
	bs := BasicScheduler{}
	bs.init()
	return &bs
}

func (s *BasicScheduler) init() {
	s.typeResPool = map[Type]*list.List{}
	s.uuidResPool = map[uuid.UUID]Resource{}
	s.requests = map[chan []Resource][]Resource{}
}

func (s *BasicScheduler) RegRes(resource Resource) util.Err {
	if s.typeResPool == nil || s.uuidResPool == nil {
		return NewUseOfNoneInitScheduler(nil)
	}
	t, id := resource.Definition()
	if id == (uuid.UUID{}) {
		if _, ok := s.typeResPool[t]; !ok {
			s.typeResPool[t] = list.New()
		}
		s.typeResPool[t].PushBack(resource)
	} else {
		s.uuidResPool[id] = resource
	}
	return nil
}

func (s *BasicScheduler) Request(ctx context.Context, logic func(), resources ...Resource) util.Err {
	//TODO: complete request function
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
		t, id := resource.Definition()
		if (id != uuid.UUID{}) {
			if res, ok := s.uuidResPool[id]; ok {
				return false, NewInvalidResourceError(nil)
			} else if res == nil {
				return false, nil
			}
		} else {
			if resList, ok := s.typeResPool[t]; !ok {
				return false, NewInvalidResourceError(nil)
			} else if resList.Len() < 1 {
				return false, nil
			}
		}
	}
	return true, nil
}

func (s *BasicScheduler) getAllRes(resources ... Resource)[]Resource{
	var gotRes []Resource
	for _, resource := range resources {
		t, id := resource.Definition()
		if (id != uuid.UUID{}) {
			gotRes = append(gotRes, s.uuidResPool[id])
			s.uuidResPool[id] = nil
		} else {
			back := s.typeResPool[t].Back()
			gotRes = append(gotRes, s.typeResPool[t].Remove(back).(Resource))
		}
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
		gotRes:=s.getAllRes(resources...)
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
			if p:=recover();p!=nil{
				success=false
			}
			delete(s.requests, ch)	//delete the request despite if request canceled
		}()
		ch<-nil
		return true
	}
	s.mutex.Lock()
	for _, resource := range resources {
		t, id := resource.Definition()
		if (id != uuid.UUID{}) {
			s.uuidResPool[id] = resource
		} else {
			s.typeResPool[t].PushBack(resource)
		}
	}
	for ch,reqRes :=range s.requests{
		if ok,_:=s.resAllAccessible(reqRes...);ok{
			if wakeUp(ch){
				ch<-s.getAllRes(reqRes...)
				break
			}
		}
	}
	s.mutex.Unlock()
}
