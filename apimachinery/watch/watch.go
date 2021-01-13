package watch

import "github.com/chauhanr/singlenetes/apimachinery/runtime"

//Interface is the watch interface that will be used for watching events as they occur.
type Interface interface {
	// Stops watching the channel and closes it
	Stop()
	// Returns a channel that receives all the events
	ResultChan() <-chan Event
}

//EventType defines possible types of events
type EventType string

// Enum type valuses for the Event type
const (
	Added    EventType = "ADDED"
	Modified EventType = "MODIFIED"
	Deleted  EventType = "DELETED"
	Error    EventType = "ERROR"
)

//Event is the structure that represents the event in s8s
type Event struct {
	Type   EventType
	Object runtime.Object
}
