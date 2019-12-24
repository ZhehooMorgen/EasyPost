package articleProvider

import "backend/util"

func NewInvalidArticleError(para interface{}) util.Err {
	err := util.NewBasicError("cannot find the article", 404, para)
	return err
}
