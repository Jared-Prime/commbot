USERNAME?=$(shell whoami)

build:
	go build -o bin/commbot cli.go

create-aws-stack:
	aws cloudformation create-stack  --stack-name $(USERNAME)-commbot-stack \
		--template-body file://.aws/cloudformation/template.yaml \
		--tags Key=creator,Value=$(USERNAME) \
		--capabilities CAPABILITY_IAM \
		--parameters file://.aws/cloudformation/parameters.json \
		--region us-east-2

update-aws-stack:
	aws cloudformation update-stack --stack-name $(USERNAME)-commbot-stack \
		--template-body file://.aws/cloudformation/template.yaml \
		--tags Key=creator,Value=$(USERNAME) \
		--capabilities CAPABILITY_IAM \
		--parameters file://.aws/cloudformation/parameters.json \
		--region us-east-2