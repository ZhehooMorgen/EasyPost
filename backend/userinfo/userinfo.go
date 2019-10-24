package userinfo

import (
	"backend/routers"
	"backend/util"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

// Start :
// Start userinfo service
func Start() error {
	http.HandleFunc(routers.Userinfo, serve)
	return nil
}

func serve(w http.ResponseWriter, req *http.Request) {
	util.CORS(w)
	var errorCode = 200
	defer func() {
		if errorCode != 200 {
			fmt.Println("Error code:", errorCode)
			w.WriteHeader(errorCode)
		} else {
			fmt.Println("Success")
		}
	}()
	fmt.Println("New request :", req.Method, req.Host, req.RequestURI)
	if req.Method != http.MethodPost {
		errorCode = 400
		return
	}
	//data : the raw []byte carrying target userID
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		errorCode = 500
		return
	}
	//userID : the final userID that we want
	userID, err := strconv.Atoi(string(data))
	if err != nil {
		errorCode = 400
		return
	}
	authHeader, ok := req.Header["Auth"]
	var auth util.Auth
	if !ok || len(authHeader) == 0 {
		auth = ""
	} else {
		auth = util.Auth(authHeader[0])
	}
	userInfo, err := getUserInfo(auth, util.UserID(userID))
	if err != nil {
		errorCode = 404
		return
	}
	str, err := json.Marshal(userInfo)
	if err != nil {
		errorCode = 500
		return
	}
	w.Write([]byte(str))
}

// get info of a user as much as the auth allows
func getUserInfo(auth util.Auth, id util.UserID) (*util.UserInfo, error) {
	var role = getRole(auth, id)
	info, err := getAllInfo(id)
	if err != nil {
		return nil, err
	}
	var ret util.UserInfo
	ret.ID = info.ID
	ret.Name = info.Name
	if role == util.Friend {
		ret.Phone = info.Phone
	}
	return &ret, nil
}

// get All info for later use
func getAllInfo(id util.UserID) (*util.UserInfo, error) {
	if id == 0 {
		return nil, ErrUserNotFound
	}
	return &util.UserInfo{
		ID:    id,
		Name:  "小明",
		Phone: "2342424fsfgg",
	}, nil
}

// get the role of requester
func getRole(auth util.Auth, targetUserID util.UserID) util.Role {
	if auth == "all" {
		return util.Friend
	}
	return util.Nobody
}
