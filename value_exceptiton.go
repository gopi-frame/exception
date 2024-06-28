package exception

import (
	"fmt"

	"github.com/gopi-frame/exception/contract"
)

// ValueException exception with value
type ValueException struct {
	contract.Throwable
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
