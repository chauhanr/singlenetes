package v1

import (
	"context"

	v1 "github.com/chauhanr/singlenetes/api/v1"
	"github.com/chauhanr/singlenetes/apimachinery/watch"
	"github.com/chauhanr/singlenetes/client-go/rest"
)

//PodsGetter interface is a window to get hte pod information.
type PodsGetter interface {
	Pods(namespace string) PodInterface
}

//PodInterface is the interface that allows us access to the Pod.
type PodInterface interface {
	Create(ctx context.Context) (*v1.Pod, error)
	Update(ctx context.Context) (*v1.Pod, error)
	List(ctx context.Context) (*v1.PodList, error)
	Watch(ctx context.Context) (watch.Interface, error)
}

type pods struct {
	client rest.Interface
	ns     string
}

func newPods(c *CoreV1Client, namespace string) *pods {
	return &pods{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

func (p *pods) Create(ctx context.Context) (*v1.Pod, error) {
	return nil, nil
}

func (p *pods) Update(ctx context.Context) (*v1.Pod, error) {
	return nil, nil
}

func (p *pods) List(ctx context.Context) (*v1.PodList, error) {
	return nil, nil
}

func (p *pods) Watch(ctx context.Context) (watch.Interface, error) {
	return nil, nil
}
