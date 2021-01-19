package cache

import (
	"github.com/chauhanr/singlenetes/apimachinery/labels"
	"github.com/chauhanr/singlenetes/apimachinery/runtime"
)

//GenericLister is an interface that lists the runtime objects
type GenericLister interface {
	List(selector string) (ret []runtime.Object, err error)
	Get(name string) (runtime.Object, error)
	ByNamespace(namespace string) GenericNamespaceLister
}

//GenericNamespaceLister a new  interface for genericLister
type GenericNamespaceLister interface {
	List(selector labels.Selector)
	Get(name string) (runtime.Object, error)
}
