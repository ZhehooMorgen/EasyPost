package util

type Err interface {
	error
	ErrorCode() int
	PreviousError() error
	ToRange() []error
}

type BasicError struct {
	error         string
	errorCode     int
	previousError error
}

func NewBasicError(error string, errorCode int, previousError error) Err {
	return &BasicError{error, errorCode, previousError}
}

func (e *BasicError) Error() string {
	return e.error
}

func (e *BasicError) ErrorCode() int {
	return e.errorCode
}

func (e *BasicError) PreviousError() error {
	return e.previousError
}

func (e *BasicError) ToRange() []error {
	var ret []error
	if e.previousError != nil {
		if err, ok := e.previousError.(Err); ok {
			ret = err.ToRange()
		}
	}
	return append(ret, e)
}

