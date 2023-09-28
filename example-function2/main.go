package main

import (
	"context"
	"fmt"

	helper "example-function2/helper"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Data string `json:"data"`
}

func HandleRequest(ctx context.Context, name MyEvent) (string, error) {
	return fmt.Sprintf("Hello %s!", helper.GetName(name)), nil
}

func main() {
	lambda.Start(HandleRequest)
}
