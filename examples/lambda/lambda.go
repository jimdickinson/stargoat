package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	stargoat "github.com/jimdickinson/stargoat/pkg"
)

type WebEvent struct {
	ID        string            `json:"id"`
	Site      string            `json:"site"`
	URL       string            `json:"url"`
	EventType string            `json:"eventType"`
	Duration  int               `json:"duration"`
	MoreData  map[string]string `json:"moreData"`
}

type Response struct {
	ID      string `json:"id"`
	Message string `json:"message"`
}

func main() {
	keyspace, ok := os.LookupEnv("ASTRA_DB_KEYSPACE")
	if !ok {
		panic("Missing ASTRA_DB_KEYSPACE env")
	}
	astraDatabaseBaseURL, ok := os.LookupEnv("ASTRA_DB_BASE_URL")
	if !ok {
		panic("Missing ASTRA_DB_BASE_URL env")
	}
	creds := &stargoat.EnvVarCredsProvider{
		UsernameEnv: "ASTRA_DB_USERNAME",
		PasswordEnv: "ASTRA_DB_PASSWORD",
	}
	token, err := stargoat.Authenticate(creds, astraDatabaseBaseURL+"/api/rest/v1/auth")
	if err != nil {
		panic(fmt.Errorf("Authentication failed: %s", err))
	}
	client, err := stargoat.NewClient(astraDatabaseBaseURL+"/api/rest", *token, true, nil)
	if err != nil {
		panic(fmt.Errorf("NewClient() failed: %s", err))
	}

	handleLambdaEvent := func(ctx context.Context, event WebEvent) (Response, error) {
		id, err := client.PutDoc(keyspace, "webevents", event.ID, event)
		if err != nil {
			return Response{Message: "ERROR"}, err
		}
		return Response{Message: "OK", ID: id}, nil
	}

	lambda.Start(handleLambdaEvent)
}
