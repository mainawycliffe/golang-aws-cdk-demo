package main

import (
	"context"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/google/uuid"
)

type Member struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func HandleRequest(ctx context.Context, member Member) (*Member, error) {
	if member.Id == "" {
		id := uuid.New()
		member.Id = id.String()
	}

	mySession := session.Must(session.NewSession())
	svc := dynamodb.New(mySession)
	item, err := dynamodbattribute.MarshalMap(member)
	if err != nil {
		return nil, err
	}
	tableName := os.Getenv("MEMBERS_TABLE_NAME")
	_, err = svc.PutItem(&dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String(tableName),
		//ReturnValues: aws.String("ALL_NEW"),
	})
	if err != nil {
		return nil, err
	}
	return &member, nil
}

func main() {
	lambda.Start(HandleRequest)
}
