package main

import (
	"testing"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/assertions"
	"github.com/aws/jsii-runtime-go"
)

func Test_upsertMemberLambdaFunction(t *testing.T) {
	app := awscdk.NewApp(nil)
	stack := NewGolangAwsCdkDemoStack(app, "TestStack", nil)

	// you can create a test table here instead
	table := MembersDynamoDBTable(stack)
	upsertMemberLambdaFunction(stack, table)

	template := assertions.Template_FromStack(stack, nil)

	template.HasResourceProperties(jsii.String("AWS::Lambda::Function"), map[string]interface{}{
		"Runtime": "go1.x",
		"Environment": map[string]interface{}{
			"Variables": map[string]interface{}{
				"MEMBERS_TABLE_NAME": map[string]interface{}{
					"Ref": assertions.Match_StringLikeRegexp(jsii.String("membersDynamoDBTable*")),
				},
			},
		},
	})

	// check for permissions
	template.HasResourceProperties(jsii.String("AWS::IAM::Policy"), map[string]interface{}{
		"PolicyDocument": map[string]interface{}{
			"Statement": []map[string]interface{}{
				{
					"Action": assertions.Match_ArrayWith(&[]interface{}{
						"dynamodb:UpdateItem",
					}),
					"Effect": "Allow",
					"Resource": assertions.Match_ArrayWith(
						&[]interface{}{
							map[string]interface{}{
								"Fn::GetAtt": []interface{}{
									assertions.Match_StringLikeRegexp(jsii.String("membersDynamoDBTable*")),
									"Arn",
								},
							},
						},
					),
				},
			},
		},
		"Roles": assertions.Match_ArrayWith(
			&[]interface{}{
				map[string]interface{}{
					"Ref": assertions.Match_StringLikeRegexp(jsii.String("upsertMemberLambdaFunctionServiceRole*")),
				},
			},
		),
	})
}
