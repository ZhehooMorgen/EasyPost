package userAction

import (
	"backend/helper/httpHelper"
	"fmt"
	"net/http"
)

func registerService(w http.ResponseWriter, req *http.Request){
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
}

type registerData struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type registerResult struct {
	UserID string `json:"userID"`
}