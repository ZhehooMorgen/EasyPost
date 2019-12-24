package userAction

import (
	"backend/helper/httpHelper"
	"backend/util"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func loginService(w http.ResponseWriter, req *http.Request) {
	httpHelper.CORS(w)
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
	//data : the raw []byte carrying {"userName":"***","password":"***"}
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		errorCode = 500
		return
	}
	var loginData loginData
	err = json.Unmarshal(data, &loginData)
	if err != nil {
		errorCode = 400
		return
	}
	fmt.Printf("%+v\n", loginData)
	var credential = `odh3fiic45t5x`
	if _, httpErr := w.Write([]byte(credential)); httpErr != nil {
		panic(httpHelper.HTTPWriteFail(httpErr))
	}
}

type loginData struct {
	UserID   util.UserID `json:"userID"`
	Password string      `json:"password"`
}

type loginResult struct {
	OK      bool   `json:"ok"`
	Session string `json:"session"`
}
