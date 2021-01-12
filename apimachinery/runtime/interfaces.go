package runtime

import "github.com/chauhanr/singlenetes/apimachinery/runtime/schema"

type Object interface {
	// object from the api machinery
	GetObjectKind() schema.GroupVersionKind
}
