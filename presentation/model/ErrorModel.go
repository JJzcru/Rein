package model

// ErrorModel Representation of an Error in presentation layer
type ErrorModel struct {
	Msg string `json:"error"`
}

// NewErrorModel Constructor for create a new ErrorModel
func NewErrorModel(msg string) ErrorModel {
	return ErrorModel{msg}
}

// GetMessage Get the error message
func (e *ErrorModel) GetMessage() string {
	return e.Msg
}
