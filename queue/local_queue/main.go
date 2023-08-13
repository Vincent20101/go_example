package main

import (
	"fmt"
	"sync"
)

var ProcedureMsgQueue = &ProcedureMsg{queue: make([]t, 0), cond: sync.NewCond(&sync.RWMutex{})}

type t interface{}

type ProcedureMsg struct {
	mu              sync.RWMutex
	queue           []t
	cond            *sync.Cond
	isEnableN40Lite bool
}

func (p *ProcedureMsg) setIsEnableN40Lite(b bool) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.isEnableN40Lite = b
}

func (p *ProcedureMsg) IsEnableN40Lite() bool {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.isEnableN40Lite
}

func (p *ProcedureMsg) Add(item interface{}) {
	p.cond.L.Lock()
	defer p.cond.L.Unlock()
	p.queue = append(p.queue, item)
	p.cond.Signal()
}

func (p *ProcedureMsg) Len() int {
	p.cond.L.Lock()
	defer p.cond.L.Unlock()
	return len(p.queue)
}

func (p *ProcedureMsg) Get() (item interface{}) {
	p.cond.L.Lock()
	defer p.cond.L.Unlock()
	for len(p.queue) == 0 {
		p.cond.Wait()
	}
	item = p.queue[0]
	p.queue[0] = nil
	p.queue = p.queue[1:]

	return
}

func main() {
	ProcedureMsgQueue.setIsEnableN40Lite(true)
	fmt.Println(ProcedureMsgQueue.IsEnableN40Lite())
}
