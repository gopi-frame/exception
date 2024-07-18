package exception

import (
	"fmt"

	ec "github.com/gopi-frame/contract/exception"
)

// ValueException exception with value
type ValueException struct {
	ec.Throwable
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
	exp.Throwable = New(message)
	return exp
}
