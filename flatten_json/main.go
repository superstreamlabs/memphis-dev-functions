package main

import (
	"encoding/json"
	"fmt"
	"github.com/memphisdev/memphis-functions.go/memphis"
)

func FlattenMessages(message []byte, headers map[string]string, inputs map[string]string) ([]byte, map[string]string, error) {
	flatten := func(out_map map[string]interface{}, value interface{}, parent_key string) {
		var recursiveFlatten func(map[string]interface{}, interface{}, string)

		recursiveFlatten = func(out_map map[string]interface{}, value interface{}, parent_key string) {
			switch value_typed := value.(type) {
			default:
				out_map[parent_key] = value
			case map[string]interface{}:
				for key, value := range value_typed {
					recursiveFlatten(out_map, value, key)
				}
			case []interface{}:
				for index, value := range value_typed {
					new_name := fmt.Sprintf("%s_%d", parent_key, index)
					recursiveFlatten(out_map, value, new_name)
				}
			}
		}

		recursiveFlatten(out_map, value, parent_key)
	}

	var msg_map map[string]interface{}
	out_struct := make(map[string]interface{})

	if err := json.Unmarshal(message, &msg_map); err != nil {
		return nil, nil, err
	}

	flatten(out_struct, msg_map, "")

	if msgStr, err := json.Marshal(out_struct); err == nil {
		return msgStr, headers, nil
	} else {
		return nil, nil, err
	}
}

func main() {
	memphis.CreateFunction(FlattenMessages)
}
