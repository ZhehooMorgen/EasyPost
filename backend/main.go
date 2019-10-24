package main

import (
	"backend/articleProvider"
	"backend/login"
	"backend/userInfo"
	"fmt"
	"net/http"
)

func main() {
	go http.ListenAndServe(":86", nil)
	fmt.Println("articleProvider service started successfully?:", login.Start() == nil)
	fmt.Println("userInfo service started successfully?:", userInfo.Start() == nil)
	fmt.Println("articleProvider service started successfully?:", articleProvider.Start() == nil)

	fmt.Println("All stuffs running fine!")
	select {}
}
