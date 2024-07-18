package exception

import ec "github.com/gopi-frame/contract/exception"

// TimeoutException timeout exception
type TimeoutException struct {
	ec.Throwable
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
