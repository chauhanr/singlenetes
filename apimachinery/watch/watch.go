package watch

import "github.com/chauhanr/singlenetes/apimachinery/runtime"

type Interface interface {
	// Stops watching the channel and closes it
	Stop()
	// Returns a channel that receives all the events
	ResultChan() <-chan Event
}

type EventType string

type Event struct {
	Type   EventType
	Object runtime.ObjectKind
}
