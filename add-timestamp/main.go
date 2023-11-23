package main

import (
	"encoding/json"
	"github.com/memphisdev/memphis-functions.go/memphis"
	"time"
)

func AddTimestamp(payload []byte, headers map[string]string, inputs map[string]string) ([]byte, map[string]string, error) {
	// Assumes JSON encoding
	var payload_json map[string]interface{}

	if err := json.Unmarshal(payload, &payload_json); err != nil {
		return nil, nil, err
	}

	payload_json["timestamp"] = time.Now().Round(time.Second)

	if modifiedPayload, err := json.Marshal(payload_json); err == nil {
		return modifiedPayload, headers, nil
	} else {
		return nil, nil, err
	}
}

func main() {
	memphis.CreateFunction(AddTimestamp)
}
