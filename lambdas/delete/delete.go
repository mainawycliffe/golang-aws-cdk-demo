package main

import (
	"context"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type DeleteMemberEvent struct {
	Id string `json:"id"`
}

func HandleRequest(ctx context.Context, member DeleteMemberEvent) (map[string]interface{}, error) {
	mySession := session.Must(session.NewSession())
	svc := dynamodb.New(mySession)
	tableName := os.Getenv("MEMBERS_TABLE_NAME")
	res, err := svc.DeleteItem(&dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(member.Id),
			},
		},
		TableName:    aws.String(tableName),
		ReturnValues: aws.String("ALL_OLD"),
	})
	if err != nil {
		return nil, err
	}
	var item map[string]interface{}
	if err := dynamodbattribute.ConvertFromMap(res.Attributes, &item); err != nil {
		return nil, err
	}
	return item, nil
}

func main() {
	lambda.Start(HandleRequest)
}
