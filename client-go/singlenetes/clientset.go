package singlenetes

import appsv1 "github.com/chauhanr/singlenetes/client-go/singlenetes/typed/apps/v1"

//Interface to the signlenetes instance.
type Interface interface {
	AppsV1() appsv1.AppsV1Interface
}
