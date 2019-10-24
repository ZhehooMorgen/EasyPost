package login

import (
	"backend/routers"
	"backend/util"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Start :
//	Start the articleProvider module
func Start() error {
	http.HandleFunc(routers.Login, serve)
	return nil
}

func serve(w http.ResponseWriter, req *http.Request) {
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
	fmt.Printf("%+v\n",loginData)
	w.Write([]byte(`odh3fiic45t5x`))
	return
}

type loginData struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}
