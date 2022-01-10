package main

import (
	"fmt"
	"github.com/sheetalpanery/golang-oauth-library/pkg/oauth"
	"net/http"
)

func main() {

	http.HandleFunc("/generateAccessToken", GenerateAccessToken)

	http.ListenAndServe(":8090", nil)

}

func GenerateAccessToken(w http.ResponseWriter, req *http.Request) {

	code := req.URL.Query()["code"]
	fmt.Println("code:", code)
	token, err := oauth.GetAccessToken(code[0])
	fmt.Println(token, err)

}
