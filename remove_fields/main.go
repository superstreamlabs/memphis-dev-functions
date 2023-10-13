package main

import (
	"encoding/json"
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

var keysToRemove[2]string

func RemoveFields(msg_str *string) ([]byte, error){
	RemoveFieldsInner := func(msgMapSubset *map[string]interface{}){
		var RecursiveRemove func(*map[string]interface{})

		// Sorta like a depth first search but it deletes the keys specified on the way out 
		RecursiveRemove = func(msgMapSubset *map[string]interface{}){
			for _, value := range *msgMapSubset{
				if  val_cast, ok := value.(map[string]interface{}); ok  {
					RecursiveRemove(&val_cast)
				}
			}

			for _, value := range keysToRemove{
				delete(*msgMapSubset, value)
			}
		}
		RecursiveRemove(msgMapSubset)	
	}
	
	var msg_map map[string]interface{}

	if err := json.Unmarshal([]byte(*msg_str), &msg_map); err != nil{
		return nil, err
	}	

	RemoveFieldsInner(&msg_map)
	
	if msgStr, err := json.Marshal(msg_map); err != nil{
		return msgStr, nil
	}else{
		return nil, err
	}
}

func RemoveFieldsHandler(ctx context.Context, event *MemphisEvent) (*MemphisEvent, error) {
	var processedEvent MemphisEvent
	for _, msg := range event.Messages {
	    msgStr := string(msg.Payload)

		msgStrRemovedFields, err := RemoveFields(&msgStr)

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
			Payload: msgStrRemovedFields,
		})
	}

	return &processedEvent, nil
}

func main() {
	keysToRemove[0] = "remove_me"
	keysToRemove[1] = "me_too"
	
	lambda.Start(RemoveFieldsHandler)
}
