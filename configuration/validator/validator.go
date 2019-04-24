package validator

// ErrorContainer represents a container where we can add errors and retrieve them
type ErrorContainer interface {
	Push(err error)
	HasErrors() bool
	Errors() []error
}

// Validator represents the validator interface
type Validator struct {
	errors []error
}

// NewValidator is a constructor of validator
func NewValidator() *Validator {
	val := new(Validator)
	val.errors = make([]error, 0)
	return val
}

// Push an error in the validator.
func (v *Validator) Push(err error) {
	v.errors = append(v.errors, err)
}

// HasErrors checks whether the validator contains errors.
func (v *Validator) HasErrors() bool {
	return len(v.errors) > 0
}

// Errors returns the errors.
func (v *Validator) Errors() []error {
	return v.errors
}
