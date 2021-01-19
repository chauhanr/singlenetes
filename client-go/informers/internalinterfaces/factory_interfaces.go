package internalinterfaces

import (
	"github.com/chauhanr/singlenetes/apimachinery/runtime"
	"github.com/chauhanr/singlenetes/client-go/tools/cache"
)

type SharedInformerFactory interface{
	Start(stopCh <-chan struct{})
	InformerFor(obj runtime.Object, newFunc NewInformerFunc) cache.SharedIndexInformer 
}

type NewInformerFunc func()