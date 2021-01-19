package singlenetes

import (
	appsv1 "github.com/chauhanr/singlenetes/client-go/singlenetes/typed/apps/v1"
	corev1 "github.com/chauhanr/singlenetes/client-go/singlenetes/typed/core/v1"
)

//Interface to the signlenetes instance.
type Interface interface {
	AppsV1() appsv1.AppsV1Interface
	CoreV1() corev1.CoreV1Interface
}
