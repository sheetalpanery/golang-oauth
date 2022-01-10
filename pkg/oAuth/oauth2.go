package oauth

import "fmt"

type oauthRequest struct {
	client_id     string
	client_secret string
	grant_type    string
	redirect_uri  string
}

func getAccessToken(code string) (token string, err error) {

	fmt.Println("code in getAccessToken", code)
	return "dummyToken", nil

}
