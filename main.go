package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/crud/dynamo"
	"github.com/crud/route"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func init() {
	dynamo.Dynamo = dynamo.ConnectDynamo()
	dynamo.CreateTable()
	err := dynamo.InjectItem()
	if err != nil {
		panic(err)
	}
	fmt.Println("connected")
}

func main() {

	func() {
		config := oauth1.NewConfig("consumerKey", "KRy7l0v8wex3w8Sy5zThai3Ea")
		token := oauth1.NewToken("accessToken", "X2eBm0Y21kYEuR74W3Frqc2JVIizOj8Q1EVGatDsEVVEJo0ucu")
		// http.Client will automatically authorize Requests
		httpClient := config.Client(oauth1.NoContext, token)

		// twitter client
		client := twitter.NewClient(httpClient)

		fmt.Println(client)

		// Home Timeline
		tweets, _, err := client.Timelines.HomeTimeline(&twitter.HomeTimelineParams{
			Count: 20,
		})

		fmt.Println(err)
		fmt.Println(tweets)

	}()

	log.Fatal(http.ListenAndServe(":8000", route.Router))

}
