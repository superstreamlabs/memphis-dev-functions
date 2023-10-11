package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func RemoveFields(keysToRemove []string, msg_map *map[string]interface{}){
	RemoveFieldsInner := func(keysToRemove []string, msgMapSubset *map[string]interface{}){
		var RecursiveRemove func([]string, *map[string]interface{})

		// Sorta like a depth first search but it deletes the keys specified on the way out 
		RecursiveRemove = func(keysToRemove []string, msgMapSubset *map[string]interface{}){
			for _, value := range *msgMapSubset{
				if  val_cast, ok := value.(map[string]interface{}); ok  {
					RecursiveRemove(keysToRemove, &val_cast)
				}
			}

			for _, value := range keysToRemove{
				delete(*msgMapSubset, value)
			}
		}
		RecursiveRemove(keysToRemove, msgMapSubset)	
	}
	
	RemoveFieldsInner(keysToRemove, msg_map)
}


func main() {
	file, err := os.ReadFile("./fieldRemovalJson.json")
	if err != nil{
		panic(err)
	}

	var msg_map map[string]interface{}
	
	if err := json.Unmarshal(file, &msg_map); err != nil{
		panic(err) // How should lambda functions fail gracefully? DLS?
	}

	var fields_to_remove []string
	fields_to_remove = append(fields_to_remove, "remove_me", "me_too")
	RemoveFields(fields_to_remove, &msg_map)

	prettyJSON, err := json.MarshalIndent(msg_map, "", "  ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(prettyJSON))
}
