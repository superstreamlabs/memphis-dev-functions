package main

import (
	"encoding/json"
	"github.com/memphisdev/memphis-functions.go/memphis"
)

func EventHandler(message []byte, headers map[string]string, inputs map[string]string) ([]byte, map[string]string,  error){
	var event map[string]interface{}
	if err := json.Unmarshal(message, &event); err != nil{
		return nil, nil, err
	}

	event[inputs["out"]] = inputs["geolocation"]
	
	if eventBytes, err := json.Marshal(event); err == nil {
		return eventBytes, headers, nil	
	} else{
		return nil, nil, err
	}
}

func main() {	
	memphis.CreateFunction(EventHandler)
}