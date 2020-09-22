package pubsub

import "sync"

type sub struct {
	ch     chan interface{}
	filter func(entry interface{}) bool
}

type pubSub struct {
	subs []*sub // 订阅者列表
	sync.RWMutex
}

func (ps *pubSub) Publish(item interface{}) {
	ps.RLock()
	defer ps.RUnlock()
	for _, sub := range ps.subs {
		if sub.filter == nil || sub.filter(item) {
			select {
			case sub.ch <- item:
			default:
			}
		}
	}
}

func (ps *pubSub) Subscribe(subChan chan interface{}, doneCh <-chan struct{}, filter func(entry interface{}) bool) {
	ps.Lock()
	defer ps.Unlock()
	s := &sub{
		ch:     subChan,
		filter: filter,
	}
	ps.subs = append(ps.subs, s)

}

func New() *pubSub {
	return &pubSub{}
}


