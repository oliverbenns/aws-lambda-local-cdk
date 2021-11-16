package main

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awslambda"
	"github.com/aws/constructs-go/constructs/v3"
	"github.com/aws/jsii-runtime-go"
)

func NewStack(scope constructs.Construct, props *awscdk.StackProps) awscdk.Stack {
	stack := awscdk.NewStack(scope, jsii.String("TestStack"), props)

	awslambda.NewFunction(stack, jsii.String("PingLambda"), &awslambda.FunctionProps{
		Code:    awslambda.Code_FromAssetImage(jsii.String("cmd/ping"), nil),
		Handler: awslambda.Handler_FROM_IMAGE(),
		Runtime: awslambda.Runtime_FROM_IMAGE(),
		Timeout: awscdk.Duration_Seconds(jsii.Number(15)),
	})

	return stack
}

func main() {
	app := awscdk.NewApp(nil)
	NewStack(app, nil)
	app.Synth(nil)
}
