package exception

import "fmt"

// ValueException exception with value
type ValueException struct {
	*Exception
	Value any
}

// NewValueException new [ValueException]
func NewValueException(v any) *ValueException {
	exp := new(ValueException)
	exp.Value = v
	message := "ValueException: Panics with value"
	if v != nil {
		message += fmt.Sprintf(": %v", v)
	}
	exp.Exception = NewException(message)
	return exp
}
