package v1

import "github.com/chauhanr/singlenetes/client-go/informers/internalinterfaces"

//Interface for Pods Informer
type Interface interface {
	Pods() PodInformer
}

type version struct {
	factory   internalinterfaces.SharedInformerFactory
	namespace string
}

//New method to ge the port informer version
func New(f internalinterfaces.SharedInformerFactory, namespace string) Interface {
	return &version{factory: f, namespace: namespace}
}

func (v *version) Pods() PodInformer {
	return &podInformer{factory: v.factory, namespace: v.namespace}
}
