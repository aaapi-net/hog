package hog

import "fmt"

// Error represents an HTTP request processing error with context.
type Error struct {
	Op      string // Operation where the error occurred
	Message string // Human-readable error description
	Err     error  // Underlying error, if any
}

// Error returns a formatted error string.
func (e *Error) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %s: %v", e.Op, e.Message, e.Err)
	}
	return fmt.Sprintf("%s: %s", e.Op, e.Message)
}

// newError creates a new Error with the specified operation, message, and underlying error.
func newError(op string, message string, err error) error {
	return &Error{
		Op:      op,
		Message: message,
		Err:     err,
	}
}
