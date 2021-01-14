package v1

import "context"

//PodsGetter interface is a window to get hte pod information.
type PodsGetter interface {
	Pods(namespace string) PodInterface
}

//PodInterface is the interface that allows us access to the Pod.
type PodInterface interface {
	Create(ctx context.Context) error
	Update(ctx context.Context) error
}
