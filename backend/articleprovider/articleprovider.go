package articleprovider

import (
	"backend/routers"
	"fmt"
	"net/http"
)

// Start :
//	Start the articleProvider module
func Start() error {
	http.HandleFunc(routers.ArticleProvider, serve)
	return nil
}

func serve(w http.ResponseWriter, req *http.Request) {
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
	str, err := getArticle(struct{}{})
	if err != nil {
		errorCode = 502
	}
	w.Write([]byte(str))
}

func getArticle(arg interface{}) (string, error) {
	if arg == nil {
		return "", nil
	}
	return "#Title \n**article**\n", nil

}
