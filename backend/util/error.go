package util

type Err interface {
	error
	ErrorCode() int
}

type BasicError struct {
	error string
	errorCode int
}

func NewBasicError(error string, errorCode int) Err {
	return &BasicError{error , errorCode}
}

func (e *BasicError)Error()string{
	return e.error
}

func (e *BasicError)ErrorCode() int {
	return e.errorCode
}