package main

import (
	"context"
	"fmt"

	helper "command-line-arguments/Users/daniel/Desktop/playground/daniel-functions/daniel-functions/function2/helper/halper.go"

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
