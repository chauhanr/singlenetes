package rest

import (
	"context"
	"io"
	"net/http"
	"net/url"

	"github.com/chauhanr/singlenetes/apimachinery/runtime"
	"github.com/chauhanr/singlenetes/apimachinery/watch"
)

//Request object for the rest client
type Request struct {
	c       *RESTClient
	verb    string
	headers http.Header
	param   url.Values

	namespace    string
	resource     string
	resourceName string
	err          error
	body         io.Reader
}

//NewRequest generated for the RESTClient
func NewRequest(c *RESTClient) *Request {

	r := &Request{
		c: c,
	}
	return r
}

//Namespace returns the Request object that can serve the api for a particular namespace
func (r *Request) Namespace(namespace string) *Request {
	return nil
}

//Resource gets the appropriate resoruce and reutrns request
func (r *Request) Resource(resource string) *Request {
	return nil
}

//Do will take a context and return a Result
func (r *Request) Do(ctx context.Context) Result {
	var result Result
	return result
}

//Verb is the method to set the verb on the request
func (r *Request) Verb(verb string) *Request {
	r.verb = verb
	return r
}

//Watch Returns a watch.Interface from the rest api
func (r *Request) Watch(ctx context.Context) (watch.Interface, error) {

	return nil, nil
}

//Result represents the return object.
type Result struct {
	body        []byte
	contentType string
	err         error
	statusCode  int
}

//Raw result in bytes returned.
func (r Result) Raw() ([]byte, error) {
	return r.body, r.err
}

//Get Returns a runtimeOnject requested
func (r Result) Get() (runtime.Object, error) {
	if r.err != nil {
		return nil, r.err
	}

	return nil, nil
}
