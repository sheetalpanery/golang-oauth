package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/generateAccessToken", generateAccessToken)

	http.ListenAndServe(":8090", nil)

}

func generateAccessToken(w http.ResponseWriter, req *http.Request) {

	code := req.URL.Query()["code"]
	fmt.Println("code:", code)
	getAccessToken

}
