package contract

import "github.com/gopi-frame/contract"

type Throwable interface {
	error
	StackTrace() string
	contract.Stringable
}
