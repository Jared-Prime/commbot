package main

import (
	"os"

	"github.com/aws/aws-lambda-go/lambda"

	commbot "github.com/jared-prime/commbot/cmd"
)

func main() {
	command := os.Getenv("LAMBDA_COMMAND_NAME")
	lambda.Start(commbot.LambdaHandler(command))
}
