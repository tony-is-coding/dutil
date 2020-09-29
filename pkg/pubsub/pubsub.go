package pubsub

import "sync"

type sub struct {
	ch     chan interface{}
	filter func(entry interface{}) bool
}

type pubSub struct {
	subs []*sub
	sync.RWMutex
}

// Publish pass item to all the filtered subscribers by channel
// Note that publish is always non-blocking so that the slow subscribe would't effect
// Therefore client need use buffered channel so as that not to miss the event
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

// Subscribe add a subscriber to the pub/sub system
func (ps *pubSub) Subscribe(subChan chan interface{}, doneCh <-chan struct{}, filter func(entry interface{}) bool) {
	ps.Lock()
	defer ps.Unlock()
	s := &sub{
		ch:     subChan,
		filter: filter,
	}
	ps.subs = append(ps.subs, s)
	go func() {
		<- doneCh
		ps.Lock()
		defer ps.Unlock()
	}()
}

func New() *pubSub {
	return &pubSub{}
}
