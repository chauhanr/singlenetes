package core

import (
	v1 "github.com/chauhanr/singlenetes/client-go/informers/core/v1"
	"github.com/chauhanr/singlenetes/client-go/informers/internalinterfaces"
)

// Interface is the interface for informers.
type Interface interface {
	V1() v1.Interface
}

type group struct {
	factory   internalinterfaces.SharedInformerFactory
	namespace string
}

//New group return function
func New(f internalinterfaces.SharedInformerFactory, namespace string) Interface {
	return &group{factory: f, namespace: namespace}
}

func (g *group) V1() v1.Interface {
	return v1.New(g.factory, g.namespace)
}
