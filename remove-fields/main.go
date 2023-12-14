package main

import (
	"encoding/json"
	"strings"

	"github.com/memphisdev/memphis-functions.go/memphis"
)

func RemoveFields(message any, headers map[string]string, inputs map[string]string) (any, map[string]string, error) {
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

	as_bytes := message.([]byte)

	if err := json.Unmarshal(as_bytes, &msg_map); err != nil {
		return nil, nil, err
	}

	RemoveFieldsInner(&msg_map)

	return msg_map, headers, nil
}

func main() {
	memphis.CreateFunction(RemoveFields)
}
