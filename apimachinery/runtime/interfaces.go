package runtime

import "github.com/chauhanr/singlenetes/apimachinery/runtime/schema"

//Object interface is a runtime interface that will represent all s8s objects.
type Object interface {
	// object from the api machinery
	GetObjectKind() schema.GroupVersionKind
}
