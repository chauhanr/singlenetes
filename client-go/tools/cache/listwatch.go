package cache

import (
	"context"

	"github.com/chauhanr/singlenetes/apimachinery/runtime"
	"github.com/chauhanr/singlenetes/apimachinery/watch"
	restclient "github.com/chauhanr/singlenetes/client-go/rest"
)

//Lister is any that knows how to perform an initial list
type Lister interface {
	List() (runtime.Object, error)
}

//Watcher is any object that knows how to start a watch on a resource.
type Watcher interface {
	Watch() (watch.Interface, error)
}

//ListWatcher is an object that can list objects and start a watcher on them
type ListWatcher interface {
	Lister
	Watcher
}

// ListFunc knows how to list resources
type ListFunc func() (runtime.Object, error)

// WatchFunc knows how to watch resources
type WatchFunc func() (watch.Interface, error)

//ListWatch structu is one for listing and watching an object
type ListWatch struct {
	ListFunc  ListFunc
	WatchFunc WatchFunc
}

//List function for the listwatch struct
func (lw *ListWatch) List() (runtime.Object, error) {
	return lw.ListFunc()
}

//Watch function for list watch struct
func (lw *ListWatch) Watch() (watch.Interface, error) {
	return lw.WatchFunc()
}

//Getter interface returns rest client reuqest object
type Getter interface {
	Get() *restclient.Request
}

//NewListWatchFromClient returns a list and watch object to load a resource on a namespace.
func NewListWatchFromClient(c Getter, resource string, namespace string) *ListWatch {
	listFunc := func() (runtime.Object, error) {
		return c.Get().Namespace(namespace).Resource(resource).Do(context.TODO()).Get()
	}
	watchFunc := func() (watch.Interface, error) {
		return c.Get().Namespace(namespace).Resource(resource).Watch(context.TODO())
	}

	return &ListWatch{listFunc, watchFunc}
}
