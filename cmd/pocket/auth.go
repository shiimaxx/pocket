package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"

	"github.com/shiimaxx/pocket"
)

func genOauthInput(consumerKey string) []byte {
	oauthInput := &pocket.OauthInput{
		ConsumerKey: consumerKey,
		RedirectURI: "http://localhost",
	}
	jsonData, _ := json.Marshal(oauthInput)
	return jsonData
}

func parseOauthOutput(res *http.Response) (*pocket.OauthOutput, error) {
	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return nil, err
	}

	var oauthOutput pocket.OauthOutput
	err = json.Unmarshal([]byte(body), &oauthOutput)
	if err != nil {
		return nil, err
	}
	return &oauthOutput, err
}

func genAuthorizeInput(consumerKey string, requestToken string) []byte {
	authorizeInput := &pocket.AuthorizeInput{
		ConsumerKey: consumerKey,
		Code:        requestToken,
	}
	jsonData, _ := json.Marshal(authorizeInput)
	return jsonData
}

func parseAuthorizeOutput(res *http.Response) (*pocket.AuthorizeOutput, error) {
	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return nil, err
	}

	var authorizeOutput pocket.AuthorizeOutput
	err = json.Unmarshal([]byte(body), &authorizeOutput)
	if err != nil {
		return nil, err
	}
	return &authorizeOutput, err
}

func Authentication(consumerKey string) (string, error) {
	uri := "https://getpocket.com/v3"
	client := &http.Client{}
	ch := make(chan struct{})

	l, err := net.Listen("tcp", "localhost:8989")
	if err != nil {
		return "", err
	}
	defer l.Close()

	go http.Serve(l, http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		ch <- struct{}{}
	}))

	// Step 2: Obtain a request token
	oauthRequestPath := "/oauth/request"
	jsonData := genOauthInput(consumerKey)

	req, err := pocket.NewRequest(uri+oauthRequestPath, jsonData)
	if err != nil {
		return "", err
	}

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}

	oauthOutput, err := parseOauthOutput(res)
	if err != nil {
		return "", err
	}

	// Step 3: Redirect user to Pocket to continue authorization
	fmt.Print("Plead access this URL. And login with your Pocket account.\n")
	fmt.Printf("https://getpocket.com/auth/authorize?request_token=%s&redirect_uri=http://localhost:8989\n", oauthOutput.Code)

	// Step 4: Receive the callback from Pocket
	<-ch

	// Step 5: Convert a request token into a Pocket access token
	authorizeRequestPath := "/oauth/authorize"
	jsonData = genAuthorizeInput(consumerKey, oauthOutput.Code)

	req, err = pocket.NewRequest(uri+authorizeRequestPath, jsonData)
	if err != nil {
		return "", err
	}

	res, err = client.Do(req)
	if err != nil {
		return "", err
	}

	authorizeOutput, err := parseAuthorizeOutput(res)
	if err != nil {
		return "", err
	}

	return authorizeOutput.AccessToken, nil
}
