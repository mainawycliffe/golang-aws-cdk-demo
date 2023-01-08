package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/jsii-runtime-go"
)

func MembersDynamoDBTable(stack awscdk.Stack) awsdynamodb.Table {
	table := awsdynamodb.NewTable(stack,
		jsii.String("membersDynamoDBTable"),
		&awsdynamodb.TableProps{
			PartitionKey: &awsdynamodb.Attribute{
				Name: jsii.String("id"),
				Type: awsdynamodb.AttributeType_STRING,
			},
			TableName:   jsii.String("Members"),
			BillingMode: awsdynamodb.BillingMode_PROVISIONED,
		},
	)
	return table
}
