package informers

import (
	"reflect"
	"sync"
	"time"
)

/*
   sharedInformerFactory is the struct that will give the informer that the client reuqires
   this will implment all the api methods that return singlenetes objects.
*/
type sharedInformerFactory struct {
	namespace     string
	lock          sync.Mutex
	defaultResync time.Duration
	customResync  map[reflect.Type]time.Duration

	informers map[reflect.Type]String
}
