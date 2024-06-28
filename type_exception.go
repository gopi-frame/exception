package exception

import "github.com/gopi-frame/exception/contract"

// TypeException type exception
type TypeException struct {
	contract.Throwable
}

func (e TypeException) Unwrap() error {
	return e.Throwable
}

// NewTypeException new type exception
func NewTypeException(message string) *TypeException {
	exp := TypeException{}
	exp.Throwable = New("TypeException: " + message)
	return &exp
}
