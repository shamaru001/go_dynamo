package dynamo

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type Person struct {
	Id          int
	Name        string
	Image       string
	Description string
}

var (
	TableName  = "people"
	RegionName = "us-west-1"
	Endpoint   = "http://localhost:3636"
	Id         = 1
)

// make new User
var User Person = Person{
	Id:          1,
	Name:        "John Doe",
	Image:       "https://www.business2community.com/wp-content/uploads/2017/08/blank-profile-picture-973460_640.png",
	Description: "Anonymous",
}

var Dynamo *dynamodb.DynamoDB

func ConnectDynamo() (db *dynamodb.DynamoDB) {
	return dynamodb.New(session.Must(session.NewSession(&aws.Config{
		Region:   &RegionName,
		Endpoint: &Endpoint,
	})))
}

func CreateTable() error {

	atrrDefinitions := []*dynamodb.AttributeDefinition{
		{
			AttributeName: aws.String("Id"),
			AttributeType: aws.String("N"),
		},
	}

	keySchema := []*dynamodb.KeySchemaElement{
		{
			AttributeName: aws.String("Id"),
			KeyType:       aws.String("HASH"),
		},
	}

	_, err := Dynamo.CreateTable(&dynamodb.CreateTableInput{
		AttributeDefinitions: atrrDefinitions,
		KeySchema:            keySchema,
		BillingMode:          aws.String(dynamodb.BillingModePayPerRequest),
		TableName:            &TableName,
	})

	return err
}

func InjectItem() error {
	return CreateItem(User)
}

