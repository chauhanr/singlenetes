package v1

import "github.com/chauhanr/singlenetes/client-go/rest"

//CoreV1Interface is the interface to constructs like pods and nodes.
type CoreV1Interface interface {
	RESTClient() rest.Interface
	PodsGetter
}
