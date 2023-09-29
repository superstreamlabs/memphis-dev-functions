package main

import (
	"context"
	"example-function2/helper"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Data string `json:"data"`
}

func HandleRequest(ctx context.Context, name MyEvent) (string, error) {
	return fmt.Sprintf("Hello %s!", helper.GetName(name.Data)), nil
}

func main() {
	lambda.Start(HandleRequest)
}
