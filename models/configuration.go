package models

import (
	"os"
	"encoding/json"
	"fmt"
	"time"
)

type Configuration struct {
	Secret string
	Pepper string
}

type BasicJSONReturn struct {
	ReturnType   string          `json:"ReturnType"        bson:"ReturnType"`
	ReturnStatus string          `json:"ReturnStatus"        bson:"ReturnStatus"`
	Payload      string          `json:"payLoad"        bson:"payLoad"`
}


func  GetSecret() string {
	file, _ := os.Open("conf/conf.json")
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println(time.Now()," Configuration.go GetSecret 001: Error: ",ErrInternalServer.Title, " ", err)
		return ""
	}
	return configuration.Secret
}

func  GetPepper() string {
	file, _ := os.Open("conf/conf.json")
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println(time.Now()," Configuration.go GetPepper 001: Error: ",ErrInternalServer.Title, " ", err)
		return ""
	}
	return configuration.Pepper
}