package v1

import "github.com/chauhanr/singlenetes/client-go/rest"

//CoreV1Interface is the interface to constructs like pods and nodes.
type CoreV1Interface interface {
	RESTClient() rest.Interface
	PodsGetter
}

//CoreV1Client object for rest interface
type CoreV1Client struct {
	restClient rest.Interface
}

//Pods function for PodsGetter
func (c *CoreV1Client) Pods(namespace string) PodInterface {
	return newPods(c, namespace)
}

//RESTClient Return the rest client to get objects
func (c *CoreV1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
