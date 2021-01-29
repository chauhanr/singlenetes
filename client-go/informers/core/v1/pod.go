package v1

import (
	"github.com/chauhanr/singlenetes/client-go/informers/internalinterfaces"
	"github.com/chauhanr/singlenetes/client-go/tools/cache"
)

//PodInformer will help in getting pod infromer interface.
type PodInformer interface {
	Informer() cache.SharedIndexInformer
	//Lister() v1.PodLister
}

type podInformer struct {
	factory   internalinterfaces.SharedInformerFactory
	namespace string
}
