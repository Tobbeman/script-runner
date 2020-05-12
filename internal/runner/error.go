package runner

import (
	"os/exec"
	"strconv"
)

type Error struct {
	// Name is the file name for which the error occurred.
	Name string
	// Err is the underlying error.
	Err error
}

func (e *Error) Error() string {
	return "exec: " + strconv.Quote(e.Name) + ": " + e.Err.Error()
}


func convertError(err error) *Error {
	if e, ok := err.(*exec.Error); ok {
		return &Error{
			e.Name,
			e.Err,
		}
	}
	return &Error{
		"Wrapped error",
		err,
	}
}