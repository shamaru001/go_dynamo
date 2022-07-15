package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/crud/dynamo"
	"github.com/crud/route"
)

func init() {
	dynamo.Dynamo = dynamo.ConnectDynamo()
	dynamo.CreateTablePeople()
	err := dynamo.InjectItem()
	if err != nil {
		panic(err)
	}

	fmt.Println("connected")
}

func main() {
	env_port := os.Getenv("PORT")
	if env_port == "" {
		env_port = "8000"
	}
	port := fmt.Sprintf(":%s", env_port)

	log.Fatal(http.ListenAndServe(port, route.Router))
}
