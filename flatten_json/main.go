package main

import (
	"encoding/json"
	"fmt"
	"context"
	"github.com/aws/aws-lambda-go/lambda"
)

type MemphisMsg struct {
	Headers map[string]string `json:"headers"`
	Payload []byte            `json:"payload"`
}

type MemphisMsgWithError struct{
	Headers map[string]string `json:"headers"`
	Payload []byte            `json:"payload"`
	Error string			  `json:"error"`
}

type MemphisEvent struct {
	Messages []MemphisMsg `json:"messages"`
	FailedMessages []MemphisMsgWithError `json:"failedMessages"`
}

func FlattenMessages(jsonStr []byte) ([]byte, error) {
	flatten := func(out_map map[string]interface{}, value interface{}, parent_key string){
		var recursiveFlatten func(map[string]interface{}, interface{}, string)

		recursiveFlatten = func(out_map map[string]interface{}, value interface{}, parent_key string){
			switch value_typed := value.(type){
			default:
				out_map[parent_key] = value
			case map[string]interface{}:
				for key, value := range value_typed{
					recursiveFlatten(out_map, value, key)
				}
			case []interface{}:
				for index, value := range value_typed{
					new_name := fmt.Sprintf("%s_%d", parent_key, index)
					recursiveFlatten(out_map, value, new_name)
				}
			}	
		}

		recursiveFlatten(out_map, value, parent_key)
	}

	var msg_map map[string]interface{}
	out_struct := make(map[string]interface{})

	if err := json.Unmarshal(jsonStr, &msg_map); err != nil{
		return nil, err
	}

	flatten(out_struct, msg_map, "")

	if msgStr, err := json.Marshal(msg_map); err != nil{
		return msgStr, nil
	}else{
		return nil, err
	}
}

func FlattenHandler(ctx context.Context, event *MemphisEvent) (*MemphisEvent, error) {
	var processedEvent MemphisEvent
	for _, msg := range event.Messages {
	    msgStr := string(msg.Payload)

		flattenedMessages, err := FlattenMessages([]byte(msgStr))

		if err != nil{
			processedEvent.FailedMessages = append(processedEvent.FailedMessages, MemphisMsgWithError{
				Headers: msg.Headers,
				Payload: []byte(msgStr),
				Error: err.Error(),
			})

			continue
		}

		processedEvent.Messages = append(processedEvent.Messages, MemphisMsg{
			Headers: msg.Headers,
			Payload: flattenedMessages,
		})
	}

	return &processedEvent, nil
}

func main() {
	lambda.Start(FlattenHandler)
}
