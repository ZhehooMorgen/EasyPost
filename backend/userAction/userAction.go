package userAction

import (
	"backend/routers"
	"backend/userAction/querys"
	"backend/util"
	"net/http"
)

// Start :
//	Start the articleProvider module
func Start() util.Err {
	http.HandleFunc(routers.UserAction+"/login", LoginService)
	if err := querys.InitMongoDBConnection(); err != nil {
		return err
	}
	return nil
}
