package cache

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
	LastSyncedResourceVersion() string
}

//SharedIndexInformer add indexes to the SharedInformer.
type SharedIndexInformer interface {
	SharedInformer

	AddIndexers(indexers Indexers)
	GetIndexer() Indexer
}

type sharedIndexInformer struct {
	indexer Indexer
}
