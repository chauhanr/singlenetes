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

//NewRESTClient returns a RESTClient that can be used to query the API server
func NewRESTClient(baseURL *url.URL, versionedAPIPath string, client *http.Client) (*RESTClient, error) {
	return &RESTClient{
		base:   baseURL,
		Client: client,
	}, nil
}

//Verb returns the appropirate request verb
func (r *RESTClient) Verb(verb string) *Request {
	return NewRequest(r).Verb(verb)
}

//Post sets up the verb as post
func (r *RESTClient) Post() *Request {
	return r.Verb("POST")
}

//Put sets up the verb as post
func (r *RESTClient) Put() *Request {
	return r.Verb("PUT")
}

//Get sets up the verb as post
func (r *RESTClient) Get() *Request {
	return r.Verb("GET")
}

//Delete sets up the verb as post
func (r *RESTClient) Delete() *Request {
	return r.Verb("DELETE")
}
