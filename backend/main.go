package main

import (
	"backend/articleprovider"
	"backend/login"
	"backend/userinfo"
	"fmt"
	"net/http"
)

func main() {
	go http.ListenAndServe(":86", nil)
	fmt.Println("articleprovider service started successfully?:", login.Start() == nil)
	fmt.Println("userinfo service started successfully?:", userinfo.Start() == nil)
	fmt.Println("articleprovider service started successfully?:", articleprovider.Start() == nil)

	fmt.Println("All stuffs running fine!")
	select {}
}
