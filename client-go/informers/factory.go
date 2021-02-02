package informers

import (
	"reflect"
	"sync"
	"time"

	"github.com/chauhanr/singlenetes/apimachinery/runtime"
	"github.com/chauhanr/singlenetes/apimachinery/runtime/schema"
	"github.com/chauhanr/singlenetes/client-go/informers/core"
	"github.com/chauhanr/singlenetes/client-go/informers/internalinterfaces"
	"github.com/chauhanr/singlenetes/client-go/singlenetes"
	"github.com/chauhanr/singlenetes/client-go/tools/cache"
)

/*
   sharedInformerFactory is the struct that will give the informer that the client reuqires
   this will implment all the api methods that return singlenetes objects.
*/
type sharedInformerFactory struct {
	client        singlenetes.Interface
	namespace     string
	lock          sync.Mutex
	defaultResync time.Duration
	customResync  map[reflect.Type]time.Duration

	informers        map[reflect.Type]cache.SharedIndexInformer
	startedInformers map[reflect.Type]bool
}

/*SharedInformerFactory Interface*/
type SharedInformerFactory interface {
	internalinterfaces.SharedInformerFactory
	ForResource(resource schema.GroupVersionKind) (GenericInformer, error)
	WaitForCacheSync(stopCh <-chan struct{}) map[reflect.Type]bool
	//Apps() apps.Interface
	Core() core.Interface
}

//NewSharedInformerFactory will return factory of shared infromers.
func NewSharedInformerFactory(client singlenetes.Interface, defaultResync time.Duration) SharedInformerFactory {
	factory := &sharedInformerFactory{
		client:           client,
		namespace:        "",
		defaultResync:    defaultResync,
		informers:        make(map[reflect.Type]cache.SharedIndexInformer),
		startedInformers: make(map[reflect.Type]bool),
		customResync:     make(map[reflect.Type]time.Duration),
	}
	return factory
}

func (f *sharedInformerFactory) Start(stopCh <-chan struct{}) {

}

func (f *sharedInformerFactory) WaitForCacheSync(stopCh <-chan struct{}) map[reflect.Type]bool {
	res := map[reflect.Type]bool{}
	return res
}

func (f *sharedInformerFactory) ForResource(resouce schema.GroupVersionKind) (GenericInformer, error) {
	return nil, nil
}

func (f *sharedInformerFactory) InformerFor(obj runtime.Object, newInfromerFunc internalinterfaces.NewInformerFunc) cache.SharedIndexInformer {
	return nil
}

func (f *sharedInformerFactory) Core() core.Interface {

	return nil
}
