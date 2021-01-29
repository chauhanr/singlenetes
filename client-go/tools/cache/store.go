package cache

//Store is an interface to define the cache store
type Store interface {
	AddObject(object interface{}) error
	Update(object interface{}) error
	Delete(object interface{}) error
	List() []interface{}
	ListKeys() []string
	Get(object interface{}) (item interface{}, exists bool, err error)
	GetByKey(key string) (item interface{}, exists bool, err error)
}

//KeyFunc knows how to make a key from the object.
type KeyFunc func(object interface{}) (string, error)

type cache struct {
	//cache storage - for thread safety
	cacheStorage ThreadSafeStore
	keyFunc      KeyFunc
}

func (c *cache) AddObject(object interface{}) error {
	return nil
}

func (c *cache) Update(object interface{}) error {
	return nil
}

func (c *cache) Delete(object interface{}) error {
	return nil
}

func (c *cache) List() []interface{} {
	return nil
}

func (c *cache) ListKeys() []string {
	return nil
}

func (c *cache) Get(object interface{}) (item interface{}, exists bool, err error) {

	return
}

func (c *cache) GetByKey(key string) (item interface{}, exists bool, err error) {
	return
}

//NewStore method retuns a cache storage back
func NewStore(keyFunc KeyFunc) Store {
	return &cache{
		cacheStorage: NewThreadSafeStore(Indexers{}, Indices{}),
		keyFunc:      keyFunc,
	}
}
