package exception

import "github.com/gopi-frame/exception/contract"

// TimeoutException timeout exception
type TimeoutException struct {
	contract.Throwable
}

func (e TimeoutException) Unwrap() error {
	return e.Throwable
}

// NewTimeoutException new timeout exception
func NewTimeoutException() *TimeoutException {
	exp := TimeoutException{
		Throwable: New("TimeoutException: timeout"),
	}
	return &exp
}
