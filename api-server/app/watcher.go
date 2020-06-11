package app

import (
	"context"
	"fmt"

	"github.com/chauhanr/singlenetes/api-server/scheme"
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
	podStream := func() <-chan scheme.PodEvent {
		eventStream := make(chan scheme.PodEvent)
		go func() {
			defer close(eventStream)
			select {
			case <-w.done:
				return
			default:
				rch := w.cli.Client.Watch(context.Background(), store.POD_KEY_PREFIX, clientv3.WithPrefix())
				for ws := range rch {
					for _, ev := range ws.Events {
						key := ev.Kv.Key
						//	value := ev.Kv.Value
						event := scheme.PodEvent{}
						event.PodDefKey = string(key)
						event.EventType = ev.Type.String()

						eventStream <- event
					}
				}
			}
		}()
		return eventStream
	}

	podStreamConsumer := func(stream <-chan scheme.PodEvent) {
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
	s := podStream()
	podStreamConsumer(s)
}

func (w *Watcher) Close() {
	close(w.done)
}
