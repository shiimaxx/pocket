package pocket

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Doer interface {
	doRequest(req *http.Request) (string, error)
}

func NewRequest(requestURL string, jsonData []byte) (*http.Request, error) {
	req, err := http.NewRequest("POST", requestURL, bytes.NewReader(jsonData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("X-Accept", "application/json")

	return req, nil
}

func NewClient(consumerKey, accessToken string) (*Client, error) {
	if len(consumerKey) == 0 {
		return nil, errors.New("Missing ConsumerKey")
	}
	if len(accessToken) == 0 {
		return nil, errors.New("Missing AccessToken")
	}
	return &Client{
		Host:        "https://getpocket.com",
		ApiEndpoint: "/v3",
		ConsumerKey: consumerKey,
		AccessToken: accessToken,
	}, nil
}

func (p *Pocketer) doRequest(req *http.Request) (string, error) {
	client := new(http.Client)
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return "", errors.New(fmt.Sprintf("HTTP status error: %s", res.StatusCode))
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

// Retrieve
// https://getpocket.com/developer/docs/v3/retrieve

func genRetrieveInput(consumerKey, accessToken string, input *RetrieveOpts) []byte {
	retrieveInput := RetrieveInput{
		ConsumerKey:  consumerKey,
		AccessToken:  accessToken,
		RetrieveOpts: *input,
	}
	jsonData, _ := json.Marshal(retrieveInput)
	return jsonData
}

func parseRetrieveOutput(body string) (*RetrieveOutput, error) {
	var retrieveOutput RetrieveOutput
	err := json.Unmarshal([]byte(body), &retrieveOutput)
	if err != nil {
		return nil, err
	}
	return &retrieveOutput, nil

}

func (c *Client) Retrieve(input *RetrieveOpts) (*RetrieveOutput, error) {
	requestURL := c.Host + c.ApiEndpoint + "/get"
	jsonData := genRetrieveInput(c.ConsumerKey, c.AccessToken, input)

	req, err := NewRequest(requestURL, jsonData)
	if err != nil {
		return nil, err
	}

	body, err := c.Pocketer.doRequest(req)
	if err != nil {
		return nil, err
	}

	return parseRetrieveOutput(body)
}

// Add
// https://getpocket.com/developer/docs/v3/add

func genAddInput(consumerKey, accessToken, url string, addOpts *AddOpts) []byte {
	addInput := AddInput{
		ConsumerKey: consumerKey,
		AccessToken: accessToken,
		URL:         url,
		AddOpts:     *addOpts,
	}
	jsonData, _ := json.Marshal(addInput)
	return jsonData
}

func parseAddOutput(body string) (*AddOutput, error) {
	var addOutput AddOutput
	err := json.Unmarshal([]byte(body), &addOutput)
	return &addOutput, err
}

func (c *Client) Add(url string, addOpts *AddOpts) (*AddOutput, error) {
	requestURL := c.Host + c.ApiEndpoint + "/add"
	jsonData := genAddInput(c.ConsumerKey, c.AccessToken, url, addOpts)

	req, err := NewRequest(requestURL, jsonData)
	if err != nil {
		return nil, err
	}

	body, err := c.Pocketer.doRequest(req)
	if err != nil {
		return nil, err
	}

	return parseAddOutput(body)
}

// Modify
// https://getpocket.com/developer/docs/v3/modify

func genModifyInput(consumerKey, accessToken string, action *Action) []byte {
	modifyInput := ModifyInput{
		ConsumerKey: consumerKey,
		AccessToken: accessToken,
		Actions:     []*Action{action},
	}
	jsonData, _ := json.Marshal(modifyInput)
	return jsonData
}

func parseModifyOutput(body string) (*ModifyOutput, error) {
	var modifyOutput ModifyOutput
	err := json.Unmarshal([]byte(body), &modifyOutput)
	if err != nil {
		return nil, err
	}
	return &modifyOutput, nil
}

func (c *Client) Modify(action *Action) (*ModifyOutput, error) {
	requestURL := c.Host + c.ApiEndpoint + "/send"
	jsonData := genModifyInput(c.ConsumerKey, c.AccessToken, action)

	req, err := NewRequest(requestURL, jsonData)
	if err != nil {
		return nil, err
	}

	body, err := c.Pocketer.doRequest(req)
	if err != nil {
		return nil, err
	}

	return parseModifyOutput(body)
}
