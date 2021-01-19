package v1

import (
	"github.com/chauhanr/singlenetes/client-go/rest"
)

//AppsV1Interface will give apps objects or api
type AppsV1Interface interface {
	RESTClient() rest.Interface
}
