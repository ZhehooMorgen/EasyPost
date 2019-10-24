package util

import "net/http"

func CORS(responseWriter http.ResponseWriter){
	responseWriter.Header().Set("Access-Control-Allow-Origin","*")
	responseWriter.Header().Set("Access-Control-Allow-Credentials","true")
}
