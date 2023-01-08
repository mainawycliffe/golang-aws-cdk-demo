package main

import (
	"testing"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/assertions"
	"github.com/aws/jsii-runtime-go"
)

// example tests. To run these tests, uncomment this file along with the
// example resource in golang-aws-cdk-demo_test.go
func TestGolangAwsCdkDemoStack(t *testing.T) {
	app := awscdk.NewApp(nil)

	stack := NewGolangAwsCdkDemoStack(app, "MyStack", nil)

	template := assertions.Template_FromStack(stack, nil)

	template.HasResourceProperties(jsii.String("AWS::Lambda::Function"), map[string]interface{}{
		"Runtime": "go1.x",
	})

	template.HasResourceProperties(jsii.String("AWS::DynamoDB::Table"), map[string]interface{}{
		"TableName": "Members",
	})
}
