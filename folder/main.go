package main

import (
	"context"
	"strings"

	"github.com/aws/aws-lambda-go/lambda"
)

type MemphisMsg struct {
	Headers map[string]string `json:"headers"`
	Payload []byte            `json:"payload"`
}

type MemphisEvent struct {
	Messages []MemphisMsg `json:"messages"`
}

func HandleRequest(ctx context.Context, event MemphisEvent) (MemphisEvent, error) {
	var processedEvent MemphisEvent
	for _, msg := range event.Messages {
		msgStr := string(msg.Payload)
		if strings.Contains(msgStr, "error") {
			msgStr = strings.Replace(msgStr, "error", "hello", -1)
		}

		processedEvent.Messages = append(processedEvent.Messages, MemphisMsg{
			Headers: msg.Headers,
			Payload: []byte(msgStr),
		})
	}

	return processedEvent, nil
}

func main() {
	lambda.Start(HandleRequest)
}
