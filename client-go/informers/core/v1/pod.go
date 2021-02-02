package v1

import (
	"context"
	"time"

	v1 "github.com/chauhanr/singlenetes/api/v1"
	"github.com/chauhanr/singlenetes/apimachinery/runtime"
	"github.com/chauhanr/singlenetes/apimachinery/watch"
	"github.com/chauhanr/singlenetes/client-go/informers/internalinterfaces"
	"github.com/chauhanr/singlenetes/client-go/singlenetes"
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

func (p *podInformer) Informer() cache.SharedIndexInformer {
	return nil
}

//PodInterface give pods in return
type PodInterface interface {
	Create(ctx context.Context, pod *v1.Pod) (*v1.Pod, error)
}

//NewPodInformer will help build the Pod Informer and use it.
func NewPodInformer(client singlenetes.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredPodInformer(client, namespace, resyncPeriod, indexers)
}

//NewFilteredPodInformer returns shared index Informer
func NewFilteredPodInformer(client singlenetes.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func() (runtime.Object, error) {
				return client.CoreV1().Pods(namespace).List(context.TODO())
			},
			WatchFunc: func() (watch.Interface, error) {
				return client.CoreV1().Pods(namespace).Watch(context.TODO())
			},
		},
		&v1.Pod{},
		resyncPeriod,
		indexers,
	)
}
