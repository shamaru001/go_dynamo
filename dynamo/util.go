package dynamo

import (
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func CreateItem(person Person) error {
	_, err := Dynamo.PutItem(&dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{
			"Id": {
				N: aws.String(strconv.Itoa(person.Id)),
			},
			"Name": {
				S: aws.String(person.Name),
			},
			"Image": {
				S: aws.String(person.Image),
			},
			"Description": {
				S: aws.String(person.Description),
			},
		},
		TableName: &PeopleTable,
	})

	return err
}

func GetItem(id int) (person Person, err error) {
	result, err := Dynamo.GetItem(&dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"Id": {
				N: aws.String(strconv.Itoa(id)),
			},
		},
		TableName: &PeopleTable,
	})

	if err != nil {
		return person, err
	}

	err = dynamodbattribute.UnmarshalMap(result.Item, &person)

	return person, err

}

func UpdateItem(person Person) error {
	_, err := Dynamo.UpdateItem(&dynamodb.UpdateItemInput{
		ExpressionAttributeNames: map[string]*string{
			"#N": aws.String("Name"),
			"#I": aws.String("Image"),
			"#D": aws.String("Description"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":Name": {
				S: aws.String(person.Name),
			},
			":Image": {
				S: aws.String(person.Image),
			},
			":Description": {
				S: aws.String(person.Description),
			},
		},
		Key: map[string]*dynamodb.AttributeValue{
			"Id": {
				N: aws.String(strconv.Itoa(person.Id)),
			},
		},
		TableName:        &PeopleTable,
		UpdateExpression: aws.String("SET #N = :Name, #I = :Image, #D = :Description"),
	})

	return err
}
