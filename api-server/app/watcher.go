package app

import (
	"context"
	"fmt"

	store "github.com/chauhanr/singlenetes/api-server/store"
	"go.etcd.io/etcd/clientv3"
)

/*
  This component will work as a watcher for various events that occur.
  Once the events occur we will push the events to the relevant consumers.
  This will be done using the callback url that the cusotmers have registered with the api-server
*/

type Watcher struct {
	cli  *store.EtcdCtl
	done chan interface{}
}

func NewWatcher(ctl *store.EtcdCtl) *Watcher {
	d := make(chan interface{})
	w := Watcher{cli: ctl, done: d}
	return &w
}

func (w *Watcher) Start() {
	eStream := func() <-chan interface{} {
		eventStream := make(chan interface{})
		go func() {
			defer close(eventStream)
			select {
			case <-w.done:
				return
			default:
				for {
					rch := w.cli.Client.Watch(context.Background(), "/api/v1/Pod", clientv3.WithPrefix())
					for ws := range rch {
						for _, ev := range ws.Events {
							eventStream <- ev
						}
					}
				}
			}
		}()
		return eventStream
	}

	streamConsumer := func(stream <-chan interface{}) {
		go func() {
			for {
				select {
				case <-w.done:
					return
				case r := <-stream:
					fmt.Printf("Event published %v\n", r)

				}
			}
		}()
	}
	s := eStream()
	streamConsumer(s)
}

func (w *Watcher) Close() {
	close(w.done)
}
