package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

// NewGolangAwsCdkDemoStack makes it easy to test the stack
func NewGolangAwsCdkDemoStack(scope constructs.Construct, id string, props *awscdk.StackProps) awscdk.Stack {
	stack := awscdk.NewStack(scope, jsii.String(id), props)

	dynamodbTable := MembersDynamoDBTable(stack)

	upsertMemberLambdaFunction(stack, dynamodbTable)
	deleteMemberLambdaFunction(stack, dynamodbTable)
	getMemberLambdaFunction(stack, dynamodbTable)

	return stack
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	NewGolangAwsCdkDemoStack(app, "GolangAwsCdkDemoStack", &awscdk.StackProps{})

	app.Synth(nil)
}
