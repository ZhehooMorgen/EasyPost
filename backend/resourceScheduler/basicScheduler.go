package resourceScheduler

import (
	"backend/util"
	"context"
	uuid "github.com/satori/go.uuid"
	"sync/atomic"
)

type BasicScheduler struct {
	typeResPool map[Type][]Resource
	uuidResPool map[uuid.UUID]Resource
	close chan struct{}
	running int32	//0 for running, 1 for not running
}

func NewBasicScheduler(){

}

func (s *BasicScheduler)init(){
	s.Close()
	s.close = make(chan struct{},1)
	s.typeResPool= map[Type][]Resource{}
	s.uuidResPool= map[uuid.UUID]Resource{}
	if atomic.CompareAndSwapInt32(&s.running,1,0){
		go s.scheduleWorker()
	}
}

func (s *BasicScheduler) Close(){
	if atomic.CompareAndSwapInt32(&s.running,0,1){
		s.close <- struct {}{}
	}
}

func (s *BasicScheduler) scheduleWorker(){
	//TODO: complete scheduleWorker
}

func (s *BasicScheduler) RegRes(resource Resource) util.Err {
	if s.typeResPool == nil || s.uuidResPool == nil {
		return NewUseOfNoneInitScheduler(nil)
	}
	t, id := resource.Definition()
	if id == (uuid.UUID{}) {
		s.typeResPool[t] = append(s.typeResPool[t], resource)
	} else {
		s.uuidResPool[id] = resource
	}
	return nil
}

func (s *BasicScheduler) Request(ctx context.Context, logic func(),resource ... Resource)util.Err{

}