package communityAction

import (
	"backend/routers"
	"backend/util"
	"net/http"
)

// Start :
// Start communityAction service
func Start() util.Err{
	http.HandleFunc(routers.CommunityAction+"/userInfo", userInfoService)
	return nil
}