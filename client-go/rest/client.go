package rest

import (
	"net/http"
	"net/url"

	"github.com/chauhanr/singlenetes/apimachinery/runtime/schema"
)

//Interface for the rest client
type Interface interface {
	//GetRateLimiter()
	Verb(verb string) *Request
	Post(verb string) *Request
	Put() *Request
	Patch() *Request
	Get() *Request
	Delete() *Request
	APIVersion() schema.GroupVersionKind
}

//RESTClient implements a unified API convention on resource paths.
type RESTClient struct {
	base           *url.URL
	versionAPIPath string
	//content config and backoff manager
	// rate limiter and warning handler
	Client *http.Client
}
