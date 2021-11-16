package main

import (
	"github.com/aws/aws-lambda-go/lambda"
)

type response struct {
	Message string
}

func handler() (*response, error) {
	res := &response{
		Message: "Pong",
	}

	return res, nil
}

func main() {
	lambda.Start(handler)
}
