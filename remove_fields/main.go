package main

import (
	"encoding/json"
	"strings"

	"github.com/memphisdev/memphis-functions.go/memphis"
)

func RemoveFields(message []byte, headers map[string]string, inputs map[string]string) ([]byte, map[string]string, error) {
	split_keys := strings.Split(inputs["keys"], ",")
	keys := make([]string, 0)
	for _, key := range split_keys {
		keys = append(keys, strings.TrimSpace(key))
	}

	RemoveFieldsInner := func(msgMapSubset *map[string]interface{}) {
		var RecursiveRemove func(*map[string]interface{})

		// Sorta like a depth first search but it deletes the keys specified on the way out
		RecursiveRemove = func(msgMapSubset *map[string]interface{}) {
			for _, value := range *msgMapSubset {
				if val_cast, ok := value.(map[string]interface{}); ok {
					RecursiveRemove(&val_cast)
				}
			}

			for _, value := range keys {
				delete(*msgMapSubset, value)
			}
		}
		RecursiveRemove(msgMapSubset)
	}

	var msg_map map[string]interface{}

	if err := json.Unmarshal(message, &msg_map); err != nil {
		return nil, nil, err
	}

	RemoveFieldsInner(&msg_map)

	if msgStr, err := json.Marshal(msg_map); err == nil {
		return msgStr, headers, nil
	} else {
		return nil, nil, err
	}
}

func main() {
	memphis.CreateFunction(RemoveFields)
}
