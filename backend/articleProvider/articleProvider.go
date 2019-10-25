package articleProvider

import (
	"backend/routers"
	"backend/util"
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
	str, err := getArticle(struct{}{})
	if err != nil {
		errorCode = err.ErrorCode()
	}
	if _, httpErr := w.Write([]byte(str)); httpErr != nil {
		panic(util.HTTPWriteFail(httpErr))
	}

}

func getArticle(arg interface{}) (string, util.Err) {
	if arg == nil {
		return "", NewInvalidArticleError(arg)
	}
	return "#Title \n**article**\n", nil
}
