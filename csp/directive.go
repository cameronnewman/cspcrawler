package csp

// Directive is a rule for a CSP directive.
type Directive interface {
	// Check the context and return whether it's allowed.
	Check(Policy, SourceContext) (bool, error)
	Get() string
}

// AllowDirective always allows access to the context.
type AllowDirective struct{}

// Check implements Directive.
func (AllowDirective) Check(Policy, SourceContext) (bool, error) {
	return true, nil
}

func (AllowDirective) Get() string {
	return ""
}
