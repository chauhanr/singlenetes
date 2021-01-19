package labels

//Labels is an interface that represents the method to work with s8s labels.
type Labels interface {
	Has(label string) (exists bool)
	Get(label string) (value string)
}
