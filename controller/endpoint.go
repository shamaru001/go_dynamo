package controller

import (
	"encoding/json"
	"net/http"

	"github.com/crud/dynamo"
)

func GetProfile(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	person, err := dynamo.GetItem(1)
	if err != nil {
		panic(err)
	}
	response, err := json.Marshal(person)
	if err != nil {
		panic(err)
	}
	res.Write(response)
}


func UpdateProfile(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	
	var user dynamo.Person
	_ = json.NewDecoder(req.Body).Decode(&user)
	
	err := dynamo.UpdateItem(user)
	if err != nil {
		panic(err)
	}
	response, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	res.Write(response)
}
