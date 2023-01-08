package main

import (
	"testing"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/assertions"
	"github.com/aws/jsii-runtime-go"
)

func TestMembersDynamoDBTable(t *testing.T) {
	type args struct {
		stack awscdk.Stack
	}

	app := awscdk.NewApp(nil)
	stack := NewGolangAwsCdkDemoStack(app, "TestStack", nil)

	MembersDynamoDBTable(stack)

	template := assertions.Template_FromStack(stack, nil)

	template.HasResourceProperties(jsii.String("AWS::DynamoDB::Table"), map[string]interface{}{
		"TableName": "Members",
		"KeySchema": []map[string]interface{}{
			{
				"AttributeName": "id",
				"KeyType":       "HASH",
			},
		},
	})
}
