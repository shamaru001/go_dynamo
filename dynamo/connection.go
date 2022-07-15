package dynamo

import (
	"os"

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
	PeopleTable = "people"
	RegionName  = "us-west-1"
	Endpoint    = os.Getenv("DYNAMODB")
	Id          = 1
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
	sess := session.Must(session.NewSession(&aws.Config{
		Region:   &RegionName,
		Endpoint: &Endpoint,
	}))
	return dynamodb.New(sess)
}

func CreateTable(atrrDefinitions []*dynamodb.AttributeDefinition, keySchema []*dynamodb.KeySchemaElement, PeopleTable string) error {
	_, err := Dynamo.CreateTable(&dynamodb.CreateTableInput{
		AttributeDefinitions: atrrDefinitions,
		KeySchema:            keySchema,
		BillingMode:          aws.String(dynamodb.BillingModePayPerRequest),
		TableName:            &PeopleTable,
	})

	if err != nil {
		panic(err)
	}

	return err
}

func InjectItem() error {
	return CreateItem(User)
}

func CreateTablePeople() error {
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

	return CreateTable(atrrDefinitions, keySchema, "people")

}

func CreateTableHistory() error {
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

	return CreateTable(atrrDefinitions, keySchema, "history")

}
