package userAction

import (
	"backend/util"
	"fmt"
	"net/http"
)

func registerService(w http.ResponseWriter, req *http.Request){
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
}
