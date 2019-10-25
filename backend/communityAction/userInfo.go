package communityAction

import (
	"backend/util"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)



func userInfoService(w http.ResponseWriter, req *http.Request) {
	util.CORS(w)
	var errorCode = 200
	defer func() {
		if errorCode != 200 {
			fmt.Println("Err code:", errorCode)
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
	if _, httpErr := w.Write([]byte(str)); httpErr != nil {
		panic(util.HTTPWriteFail(httpErr))
	}
}

// get info of a user as much as the auth allows
func getUserInfo(auth util.Auth, id util.UserID) (*UserInfo, error) {
	var role = getRole(auth, id)
	info, err := getAllInfo(id)
	if err != nil {
		return nil, err
	}
	var ret UserInfo
	ret.ID = info.ID
	ret.Name = info.Name
	if role == util.Friend {
		ret.Phone = info.Phone
	}
	return &ret, nil
}

// get All info for later use
func getAllInfo(id util.UserID) (*UserInfo, error) {
	if id == 0 {
		return nil, ErrUserNotFound
	}
	return &UserInfo{
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

//UserInfo : the data structure to describe user
type UserInfo struct {
	//ID : the unique uint64 integer to identify user
	ID util.UserID `json:"id"`
	//Name : the users display name, can be changed frequently
	Name string `json:"userName"`
	//password : should never get accessible
	password string
	//Phone ：phone number or other contact method of a user, can not leak
	Phone string `json:"phone"`
}