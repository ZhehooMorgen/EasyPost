package util

type Error interface {
	error
	ErrorCode() int
}

type BasicError struct {
	error string
	errorCode int
}

func NewBasicError(error string, errorCode int) Error{
	return &BasicError{error , errorCode}
}

func (e *BasicError)Error()string{
	return e.error
}

func (e *BasicError)ErrorCode() int {
	return e.errorCode
}