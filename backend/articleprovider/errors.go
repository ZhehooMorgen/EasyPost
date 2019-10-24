package articleprovider

import "backend/util"

type InvalidArticleError struct {
	util.Error
	para interface{} //what para provided to get article
}

func NewInvalidArticleError(para interface{}) util.Error {
	return InvalidArticleError{util.NewBasicError("cannot find the article", 404), para}
}
