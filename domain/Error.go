package domain

// Error Representation of an Error
type Error struct {
	code    int
	message string
}

// NewError Constructor for create a new Error
func NewError(code int, msg string) Error {
	return Error{code, msg}
}

// GetMessage Get the error message
func (e *Error) GetMessage() string {
	return e.message
}
