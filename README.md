# Example Secure Counter Server

Example server for [amazon-cognito-identity-dart](https://github.com/jonsaw/amazon-cognito-identity-dart/tree/master/example).

## Getting Started

1. Clone repository.

	```
	git clone git@github.com:jonsaw/example-secure-counter-server.git
	```

1. Setup DynamoDB Table.

	```
	Properties:
	  TableName: counter-dev
	  AttributeDefinitions:
	    - AttributeName: userId
	      AttributeType: S
	  KeySchema:
	    - AttributeName: userId
	        KeyType: HASH
	  ProvisionedThroughput:
	    ReadCapacityUnits: 1
	    WriteCapacityUnits: 1
	```
1. Get dependencies and build.

	```
	make
	```

1. Install [serverless](https://serverless.com/).

	```
	npm install serverless -g
	```

1. Deploy & take note of API Endpoint Id.

	```
	serverless deploy
	```

1. Setup AWS Cognito.

1. Give AWS IAM Role access to Lambda function. Replace xxxxxxxxxx with Endpoint Id.

	```json
	{
	  "Version": "2012-10-17",
	  "Statement": [
	    {
	      "Effect": "Allow",
	      "Action": [
	        "mobileanalytics:PutEvents",
	        "cognito-sync:*",
	        "cognito-identity:*"
	      ],
	      "Resource": [
	        "*"
	      ]
	    },
	    {
	      "Effect": "Allow",
	      "Action": [
	        "execute-api:Invoke"
	      ],
	      "Resource": [
	        "arn:aws:execute-api:ap-southeast-1:*:xxxxxxxxxx/*"
	      ]
	    }
	  ]
	}
	```
