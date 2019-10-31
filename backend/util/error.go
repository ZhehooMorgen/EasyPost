package util

/*
	Error code rule
	Use http error code as much as possible
	Unknown error use sub zero error code
	Error code under 100 represent system and program error:
		10: Program logic unrecoverable error
		20: Network to specific service unrecoverable disconnected
		74: Fail due to context is canceled
*/

//Design:
//	Err is a linked list end with an interface,
//	That interface can be
type Err interface {
	error
	ErrorCode() int
	Previous() interface{}
	ToRange() []interface{}
}

type BasicError struct {
	error     string
	errorCode int
	previous  interface{}
}

func NewBasicError(error string, errorCode int, previous interface{}) Err {
	return &BasicError{error, errorCode, previous}
}

func (e *BasicError) Error() string {
	return e.error
}

func (e *BasicError) ErrorCode() int {
	return e.errorCode
}

func (e *BasicError) Previous() interface{} {
	return e.previous
}

func (e *BasicError) ToRange() []interface{} {
	var ret []interface{}
	if e.previous != nil {
		if err, ok := e.previous.(Err); ok {
			ret = err.ToRange()
		}else{
			ret = append(ret, e.previous)
		}
	}
	return append(ret, e)
}

func NewContextCanceled( previousError interface{}) Err {
	return NewBasicError("Context is canceled, return",74,previousError)
}

