package informers

import (
	"reflect"
	"sync"
	"time"

	"github.com/chauhanr/singlenetes/client-go/tools/cache"
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

	informers        map[reflect.Type]cache.SharedIndexInformer
	startedInformers map[reflect.Type]bool
}
