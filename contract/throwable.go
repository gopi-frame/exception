package contract

import "github.com/gopi-frame/support/contract"

type Throwable interface {
	error
	StackTrace() string
	contract.Stringable
}
