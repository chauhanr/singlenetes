package v1

//EventsGetter helps us get handle on the Event Interface
type EventsGetter interface {
	Events(namespace string) EventInterface
}

//EventInterface is the one that helps get events infromation
type EventInterface interface {
}
