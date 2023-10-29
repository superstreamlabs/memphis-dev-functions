package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/memphisdev/memphis.go"
)

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

func main() {
	lambda.Start(memphis.CreateFunction(FlattenMessages))
}