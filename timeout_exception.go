package exception

// TimeoutException timeout exception
type TimeoutException struct {
	*Exception
}

// NewTimeoutException new timeout exception
func NewTimeoutException() *TimeoutException {
	return &TimeoutException{
		Exception: NewException("TimeoutException: timeout"),
	}
}
