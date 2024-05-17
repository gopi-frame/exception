package exception

import "fmt"

// RangeException range exception
type RangeException struct {
	*Exception
	start *int
	end   *int
}

func (e *RangeException) Error() string {
	message := "RangeException"
	if e.start == nil {
		if e.end != nil {
			message += fmt.Sprintf(": index can not be greater than %d", *e.end)
		}
	} else if e.end == nil {
		message += fmt.Sprintf(": index can not be less than %d", *e.start)
	} else if *e.end > *e.start {
		message += fmt.Sprintf(": index is not in inclusive range %d..%d", *e.start, *e.end)
	} else if *e.end < *e.start {
		message += ": valid index value is empty"
	} else {
		message += fmt.Sprintf(": index must be %d", *e.start)
	}
	return message
}

// NewRangeException new range exception
func NewRangeException(start int, end int) *RangeException {
	exp := new(RangeException)
	exp.start = &start
	exp.end = &end
	exp.Exception = NewException()
	return exp
}
