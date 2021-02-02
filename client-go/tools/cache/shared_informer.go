package cache

import (
	"sync"
	"time"

	"github.com/chauhanr/singlenetes/apimachinery/runtime"
)

/*SharedInformer provides an eventual consistent linkage between clients to the authoritative state
  of a given collection of singlenetes objects.*/
type SharedInformer interface {
	AddEventHandler()
	GetStore() Store
	// GetController Controller
	/*Run starts the informer, returning when it stops i.e. when stopCh ends*/
	Run(stopCh <-chan struct{})
	/*this returns true when the informer has reached a authritative state.*/
	HasSynced() bool
	/*LastSyncedResourceVersion returns the version of the resource that was returned when
	  last sync happened.*/
	LastSyncResourceVersion() string
}

//SharedIndexInformer add indexes to the SharedInformer.
type SharedIndexInformer interface {
	SharedInformer

	AddIndexers(indexers Indexers)
	GetIndexer() Indexer
}

type sharedIndexInformer struct {
	indexer       Indexer
	listerWatcher ListWatcher
	objectType    runtime.Object

	resyncedPeriod                  time.Duration
	defaultEventHandlerResyncPeriod time.Duration
	started, stopped                bool
	startedLock                     sync.Mutex
	blockedDeltas                   sync.Mutex
}

func (s *sharedIndexInformer) AddEventHandler() {

}

func (s *sharedIndexInformer) GetStore() Store {
	return nil
}

func (s *sharedIndexInformer) Run(stopCh <-chan struct{}) {

}

func (s *sharedIndexInformer) HasSynced() bool {
	return false
}

func (s *sharedIndexInformer) LastSyncResourceVersion() string {
	return ""
}

func (s *sharedIndexInformer) GetIndexer() Indexer {
	return s.indexer
}

func (s *sharedIndexInformer) AddIndexers(indexers Indexers) {
	return
}

//NewSharedIndexInformer returns index informer
func NewSharedIndexInformer(lw ListWatcher, object runtime.Object, defaultEventHandlerResyncPeriod time.Duration, indexers Indexers) SharedIndexInformer {
	sIndexInformer := &sharedIndexInformer{
		listerWatcher:  lw,
		objectType:     object,
		resyncedPeriod: defaultEventHandlerResyncPeriod,
		indexer:        NewIndexer(indexers),
	}
	return sIndexInformer
}
