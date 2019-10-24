package articleProvider

import "backend/util"

type InvalidArticleError struct {
	util.Err
	para interface{} //what para provided to get article
}

func NewInvalidArticleError(para interface{}) util.Err {
	err:=InvalidArticleError{util.NewBasicError("cannot find the article", 404), para}
	err.Error()
	return err
}
