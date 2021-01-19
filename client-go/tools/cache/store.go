package cache

type Store interface {
	AddObject(object interface{}) error
	Update(object interface{}) error
	Delete(object interface{}) error
	List() []interface{}
	ListKeys() []string
	Get(object interface{}) (item interface{}, exists bool, err error)
	GetByKey(key string) (item interface{}, exists bool, err error)
}
