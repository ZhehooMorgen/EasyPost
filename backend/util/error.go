package util

/*
	Error code rule
	Use http error code as much as possible
	Unknown error use sub zero error code
	Error code under 100 represent system and program error:
		10: Program logic unrecoverable error
		20: Network to specific service unrecoverable disconnected
*/


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

