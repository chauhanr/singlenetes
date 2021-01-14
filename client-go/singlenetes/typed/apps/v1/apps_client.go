package v1

import "github.com/chauhanr/singlenetes/client-go/rest"

//AppsV1Interface is the interface that returns api
type AppsV1Interface interface {
	RESTClient() rest.Interface
}
