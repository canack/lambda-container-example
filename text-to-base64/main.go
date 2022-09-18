package main

import (
	"context"
	"encoding/base64"
	"github.com/aws/aws-lambda-go/lambda"
)

type Request struct {
	Text string `json:"text"`
}

func textToBase64(text string) string {
	return base64.StdEncoding.EncodeToString([]byte(text))
}

func HandleRequest(ctx context.Context, request Request) (string, error) {
	b64 := textToBase64(request.Text)
	return b64, nil
}

func main() {
	lambda.Start(HandleRequest)
}
