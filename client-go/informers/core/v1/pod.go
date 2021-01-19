package v1

import "github.com/chauhanr/singlenetes/client-go/tools/cache"

type PodInformer interface {
	Informer() cache.SharedIndexInformer
	//Lister() v1.PodLister
}