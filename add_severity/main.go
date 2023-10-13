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

func CheckSeverity(jsonStr *string) ([]byte, error){
	var msg_map map[string]interface{}

	if err := json.Unmarshal([]byte(*jsonStr), &msg_map); err != nil{
		return nil, err
	}	

	last_produce_from_message_ms, ok := ((msg_map)["time_since_last_produce"]).(float64)
	last_produce_from_message_ms_int := int(last_produce_from_message_ms)
	if !ok{
		fmt.Println("last_produce_from_message is not an int", (msg_map)["time_since_last_produce"])
		err_str := "time_since_last_produce is not a valid key, or was not parsed correctly as a fload by json.Unmarshal"
		return nil, fmt.Errorf(err_str)
	}

	if last_produce_from_message_ms_int >= 10_000{
		(msg_map)["severity"] = "critical"
	}else if last_produce_from_message_ms_int >= 5000{
		(msg_map)["severity"] = "high" 
	} else{
		(msg_map)["severity"] = "low"
	}

	if msgStr, err := json.Marshal(msg_map); err != nil{
		return msgStr, nil
	}else{
		return nil, err
	}
}

func CheckSeverityHandler(ctx context.Context, event *MemphisEvent) (*MemphisEvent, error) {
	var processedEvent MemphisEvent
	for _, msg := range event.Messages {
	    msgStr := string(msg.Payload)

		severityMsg, err := CheckSeverity(&msgStr)

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
			Payload: severityMsg,
		})
	}

	return &processedEvent, nil
}

func main() {
	lambda.Start(CheckSeverity)
}
