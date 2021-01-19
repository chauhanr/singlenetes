package labels

import "github.com/chauhanr/singlenetes/apimachinery/selection"

//Selector represents label selector
type Selector interface {
	Matches(Labels) bool
	// returns true if the selector does not restrict the selection space
	Empty() bool
	// String returns human readable string that represents selector
	String() string
	// Add requirements
	Add(r ...Requirement) Selector

	Requirements() (requirements Requirement, selectorable bool)

	DeepCopySelector() Selector

	RequirementsExactMatch(label string) (value string, found bool)
}

//Requirement contains a values, key and opertor that relates the key and values.
type Requirement struct {
	key      string
	operator selection.Operator

	strValues []string
}
