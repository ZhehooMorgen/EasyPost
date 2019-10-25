package userAction

import (
	"backend/routers"
	"backend/userAction/dataBase"
	"backend/util"
	"net/http"
)

// Start :
//	Start the articleProvider module
func Start() util.Err {
	http.HandleFunc(routers.UserAction+"/login", LoginService)
	if err := dataBase.ConnectMongoDB(); err != nil {
		return err
	}
	return nil
}
