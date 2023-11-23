package main

import (
	"encoding/json"
	"github.com/memphisdev/memphis-functions.go/memphis"
	"time"
)

type IntConversionError struct {
	message string
}

func (e *IntConversionError) Error() string {
	return e.message
}

func UnixToDateTime(payload []byte, headers map[string]string, input map[string]string) ([]byte, map[string]string, error) {
	// Assumes JSON encoding
	var payload_json map[string]interface{}

	if err := json.Unmarshal(payload, &payload_json); err != nil {
		return nil, nil, err
	}

	// Assumes input["timestamp"] is the name of the timestamp field in your payload
	unix_time := payload_json[input["timestamp"]]

	// Type Assertion
	if unix_time_64, ok := unix_time.(int64); ok {
		payload_json[input["out"]] = time.Unix(unix_time_64, 0).Round(time.Second)
	} else {
		return nil, nil, &IntConversionError{message: "key input['timestamp'] returned a value that could not be convereted into an int64"}
	}

	if modifiedPayload, err := json.Marshal(payload_json); err == nil {
		return modifiedPayload, headers, nil
	} else {
		return nil, nil, err
	}
}

func main() {
	memphis.CreateFunction(UnixToDateTime)
}
