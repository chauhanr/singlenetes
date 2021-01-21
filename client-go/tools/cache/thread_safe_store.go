package cache

import "sync"

//ThreadSafeStore is the inteface that for the thread safe store that will store the keys and index them.
type ThreadSafeStore interface {
	Add(key string, object interface{})
	Update(key string, object interface{})
	Delete(key string)
	Get(key string) (item interface{}, exists bool)
	List() []interface{}
	ListKeys() []string
	Replace(map[string]interface{}, string)
	Index(indexName string, object interface{}) ([]interface{}, error)
	GetIndexers() Indexers
	// do not add indexers after you have added data to the store.
	AddIndexers(newIndexers Indexers) error
}

type threadSafeMap struct {
	lock  sync.RWMutex
	items map[string]interface{}

	indexers Indexers
	indices  Indices
}

func (t *threadSafeMap) Add(key string, object interface{}) {
	return
}

func (t *threadSafeMap) Update(key string, object interface{}) {
	return
}

func (t *threadSafeMap) Delete(key string) {
	return
}

func (t *threadSafeMap) Get(key string) (items interface{}, exists bool) {
	return
}

func (t *threadSafeMap) List() []interface{} {
	return nil
}

func (t *threadSafeMap) ListKeys() []string {

	return nil
}

func (t *threadSafeMap) Replace(o map[string]interface{}, v string) {

	return
}

func (t *threadSafeMap) Index(indexName string, object interface{}) ([]interface{}, error) {
	return nil, nil
}

func (t *threadSafeMap) GetIndexers() Indexers {
	return nil
}

func (t *threadSafeMap) AddIndexers(newIndexers Indexers) error {
	return nil
}

//NewThreadSafeStore returns a version of a threadsafe storage
func NewThreadSafeStore(indexers Indexers, indices Indices) ThreadSafeStore {
	return &threadSafeMap{
		items:    map[string]interface{}{},
		indexers: indexers,
		indices:  indices,
	}
}
