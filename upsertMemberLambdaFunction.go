package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdklambdagoalpha/v2"
	"github.com/aws/jsii-runtime-go"
)

func upsertMemberLambdaFunction(stack awscdk.Stack, membersTable awsdynamodb.Table) awslambda.Function {
	function := awscdklambdagoalpha.NewGoFunction(stack,
		jsii.String("upsertMemberLambdaFunction"),
		&awscdklambdagoalpha.GoFunctionProps{
			Entry:   jsii.String("./lambdas/upsert"),
			Runtime: awslambda.Runtime_GO_1_X(),
			Environment: &map[string]*string{
				// Pass the table name to the lambda function
				"MEMBERS_TABLE_NAME": membersTable.TableName(),
			},
		},
	)

	// Grant the lambda function read/write permissions to our table
	membersTable.GrantReadWriteData(function)
	return function
}
