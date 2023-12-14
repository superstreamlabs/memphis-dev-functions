package main

import (
	"encoding/json"
	"github.com/memphisdev/memphis-functions.go/memphis"
	"time"
)

func AddTimestamp(payload any, headers map[string]string, inputs map[string]string) (any, map[string]string, error) {
	// Assumes JSON encoding
	as_bytes := payload.([]byte)
	var payload_json map[string]interface{}

	if err := json.Unmarshal(as_bytes, &payload_json); err != nil {
		return nil, nil, err
	}

	payload_json["timestamp"] = time.Now().Round(time.Second)

	return payload_json, headers, nil
}

func main() {
	memphis.CreateFunction(AddTimestamp)
}
