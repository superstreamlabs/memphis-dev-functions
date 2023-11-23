package main

import (
	"encoding/json"
	"strconv"

	"github.com/memphisdev/memphis-functions.go/memphis"
)

type ConversionError struct {
	message string
}

func (e *ConversionError) Error() string {
	return e.message
}

func CheckSeverity(message []byte, headers map[string]string, inputs map[string]string) ([]byte, map[string]string, error) {
	var msgMap map[string]interface{}

	if err := json.Unmarshal([]byte(message), &msgMap); err != nil {
		return nil, nil, err
	}

	var measuredValue float64

	if msgField, ok := msgMap[inputs["field"]].(float64); ok {
		measuredValue = msgField
	} else {
		return nil, nil, &ConversionError{message: "Given field key was not able to be converted to a float"}
	}

	severityCutoff, err := strconv.ParseFloat(inputs["cutoff"], 32)

	if err != nil {
		return nil, nil, err
	}

	if measuredValue >= severityCutoff {
		(msgMap)["severity"] = inputs["high"]
	} else {
		(msgMap)["severity"] = inputs["low"]
	}

	if msgStr, err := json.Marshal(msgMap); err == nil {
		return msgStr, headers, nil
	} else {
		return nil, nil, err
	}
}

func main() {
	memphis.CreateFunction(CheckSeverity)
}
