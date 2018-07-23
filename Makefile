include .env

USERNAME?=$(shell whoami)

build:
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o .aws/lambda/commbot cli.go

zip-lambda:
	zip lambda.zip .aws/lambda/commbot

upload-lambda:
	aws s3 cp lambda.zip s3://jared-prime

create-aws-stack: # build zip-lambda upload-lambda
	aws cloudformation create-stack  --stack-name $(USERNAME)-commbot-stack \
		--template-body file://.aws/cloudformation/template.yaml \
		--tags Key=creator,Value=$(USERNAME) \
		--capabilities CAPABILITY_IAM \
		--parameters file://.aws/cloudformation/parameters.json \
		--region us-east-2