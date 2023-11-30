package helper

import (
	"email-service/model"
	"encoding/json"
	"log"
)

func AssertJsonToStruct(body []byte) model.UserCredential {
	var credential model.UserCredential
	if err := json.Unmarshal(body, &credential); err != nil {
		log.Fatal(err)
	}
	return credential
}