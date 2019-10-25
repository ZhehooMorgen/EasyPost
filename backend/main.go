package main

import (
	"backend/articleProvider"
	"backend/communityAction"
	"backend/userAction"
	"fmt"
	"net/http"
)

func main() {
	go http.ListenAndServe(":86", nil)
	fmt.Println("articleProvider service started successfully?:", userAction.Start() == nil)
	fmt.Println("communityAction service started successfully?:", communityAction.Start() == nil)
	fmt.Println("articleProvider service started successfully?:", articleProvider.Start() == nil)

	fmt.Println("All stuffs running fine!")
	select {}
}
