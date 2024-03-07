package helper

// Helper Error is a custom error type that implements the error interface.
type Error struct {
	Code    int    // Status code
	Message string // Error message
}

// Error implements the error interface
func (e *Error) Error() string {
	return e.Message
}

func (e *Error) StatusCode() int {
	return e.Code
}
