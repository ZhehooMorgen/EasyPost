package httpHelper

import "backend/util"

func HTTPWriteFail(err error) util.Err {
	return util.NewBasicError("failed to response http request", 20, err)
}

func HTTPMethodWrong(err error) util.Err {
	return util.NewBasicError("the http request has wrong method", 405, err)
}

func HTTPReadFail(err error) util.Err {
	return util.NewBasicError("can not read request data", 500, err)
}

func ReqDataParseFail(causedByServerCode bool, err error) util.Err {
	var errorCode int
	var errorMessage ="can not parse data from request"
	if causedByServerCode {
		errorCode = 500
		errorMessage+=", probably caused by parse target is nil or not a pointer"
	} else {
		errorCode = 400
	}
	return util.NewBasicError(errorMessage, errorCode, err)
}
