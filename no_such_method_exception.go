package exception

import (
	"fmt"
	"reflect"

	ec "github.com/gopi-frame/contract/exception"
)

// NoSuchMethodException no such method exception
type NoSuchMethodException struct {
	ec.Throwable
	method       string
	receiverType reflect.Type
}

func (e NoSuchMethodException) Unwrap() error {
	return e.Throwable
}

// NewNoSuchMethodException new no such method exception
func NewNoSuchMethodException(receiver any, method string) *NoSuchMethodException {
	exp := new(NoSuchMethodException)
	exp.method = method
	exp.receiverType = reflect.TypeOf(receiver)
	exp.Throwable = New(fmt.Sprintf("NoSuchMethodException: no such method %s in %s", exp.method, exp.receiverType.Name()))
	return exp
}

// NewUnexportedMethodException new unexport method exception
func NewUnexportedMethodException(method string) *NoSuchMethodException {
	exp := NoSuchMethodException{}
	exp.method = method
	exp.Throwable = New(fmt.Sprintf("NoSuchMethodException: method %s is not exported", method))
	return &exp
}
