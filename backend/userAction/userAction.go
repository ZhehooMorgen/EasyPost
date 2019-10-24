package userAction

import (
	"backend/routers"
	"backend/util"
	"net/http"
)

// Start :
//	Start the articleProvider module
func Start() util.Err {
	http.HandleFunc(routers.UserAction+"/login", LoginService)
	return nil
}