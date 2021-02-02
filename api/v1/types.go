package v1

import "github.com/chauhanr/singlenetes/apimachinery/runtime/schema"

//Pod Struct is defined here
type Pod struct {
}

//GetObjectKind for Pod
func (p *Pod) GetObjectKind() schema.GroupVersionKind {
	gvk := schema.GroupVersionKind{}
	return gvk
}

//PodList gives a list of pods
type PodList struct {
}

/*GroupVersionKind return ObjectKind respresentation of pod list
func (p *PodList) GroupVersionKind() schema.GroupVersionKind {
	gvk := schema.GroupVersionKind{}
	return gvk
}

//SetGroupVersionKind method sets gvk
func (p *PodList) SetGroupVersionKind(gvk schema.GroupVersionKind) {
	return
}*/

//GetObjectKind returns ObjectKind
func (p *PodList) GetObjectKind() schema.GroupVersionKind {
	gvk := schema.GroupVersionKind{}
	return gvk

}
