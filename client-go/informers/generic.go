package informers

import "github.com/chauhanr/singlenetes/client-go/tools/cache"

type GenericInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() cache.GenericLister
}
