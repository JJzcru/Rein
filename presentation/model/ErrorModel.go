package model

// ErrorModel Representation of an Error in presentation layer
type ErrorModel struct {
	code    int
	message string
}

// NewErrorModel Constructor for create a new ErrorModel
func NewErrorModel(code int, msg string) ErrorModel {
	return ErrorModel{code, msg}
}

// GetMessage Get the error message
func (e *ErrorModel) GetMessage() string {
	return e.message
}
