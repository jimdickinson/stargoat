package stargoat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type StargateToken struct {
	AuthToken string `json:"authToken"`
}

func Authenticate(creds CredsProvider, authURL string) (*StargateToken, error) {
	type authbody struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	username, password, err := creds.Get()
	if err != nil {
		return nil, err
	}
	jsonBytes, err := json.Marshal(authbody{Username: username, Password: password})
	if err != nil {
		return nil, err
	}
	body := bytes.NewReader(jsonBytes)
	request, err := http.NewRequest("POST", authURL, body)
	if err != nil {
		return nil, err
	}
	request.Header.Add("Content-Type", "application/json")
	// TODO provide a way to use a client that's not the default
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	if response.StatusCode >= 300 {
		log.Printf("%s", string(responseBytes))
		return nil, fmt.Errorf("authentication failed, %d status code", response.StatusCode)
	}
	var responseToken StargateToken
	err = json.Unmarshal(responseBytes, &responseToken)
	if err != nil {
		return nil, err
	}
	return &responseToken, nil
}
