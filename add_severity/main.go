package main

import (
	"encoding/json"
	"os"
	"fmt"
)

func CheckSeverity(msg_map *map[string]interface{}){
	last_produce_from_message_ms, ok := ((*msg_map)["time_since_last_produce"]).(float64)
	last_produce_from_message_ms_int := int(last_produce_from_message_ms)
	if !ok{
		fmt.Println("last_produce_from_message is not an int", (*msg_map)["time_since_last_produce"])
		return //Maybe change headers to send to DLS?
	}
	if last_produce_from_message_ms_int >= 10_000{
		(*msg_map)["severity"] = "critical"
	}else if last_produce_from_message_ms_int >= 5000{
		(*msg_map)["severity"] = "high" 
	} else{
		(*msg_map)["severity"] = "low"
	}
}


func main() {
	file, err := os.ReadFile("./severity.json")
	if err != nil{
		panic(err)
	}

	var msg_map map[string]interface{}
	
	if err := json.Unmarshal(file, &msg_map); err != nil{
		panic(err) // How should lambda functions fail gracefully? DLS?
	}

	CheckSeverity(&msg_map)

	fmt.Println(msg_map)
}
