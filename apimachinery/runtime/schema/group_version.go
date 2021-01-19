package schema

//GroupVersionKind for GVK
type GroupVersionKind struct {
	Group   string
	Version string
	Kind    string
}

//GroupVersionResource for GVR
type GroupVersionResource struct {
	Group    string
	Version  string
	Resource string
}
