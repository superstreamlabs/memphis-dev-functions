package main

import (
	"encoding/json"
	"github.com/memphisdev/memphis-functions.go/memphis"
	"net/http"
	"io"
	"fmt"
)

func EventHandler(message []byte, headers map[string]string, inputs map[string]string) ([]byte, map[string]string,  error){
	var event map[string]interface{}
	if err := json.Unmarshal(message, &event); err != nil{
		return nil, nil, err
	}

	url := fmt.Sprintf("http://ip-api.com/json/%s", inputs["geolocation"])
	res, err := http.Get(url)

	if err != nil{
		return nil, nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	
	if err != nil{
		return nil, nil, err
	}

	event[inputs["out"]] = string(body) 
	
	if eventBytes, err := json.Marshal(event); err == nil {
		return eventBytes, headers, nil	
	} else{
		return nil, nil, err
	}
}

func main() {	
	memphis.CreateFunction(EventHandler)
}