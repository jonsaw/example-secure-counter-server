package main

import (
	"counter/models"
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

const region = "ap-southeast-1"
const counterTableName = "counter-dev"

// Handler handles read counter requests
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
		}, err
	}

	svc := dynamodb.New(sess)
	input := &dynamodb.UpdateItemInput{
		TableName: aws.String(counterTableName),
		Key: map[string]*dynamodb.AttributeValue{
			"userId": {
				S: aws.String(request.RequestContext.Identity.CognitoIdentityID),
			},
		},
		UpdateExpression: aws.String("SET #count = if_not_exists(#count, :start) + :num"),
		ExpressionAttributeNames: map[string]*string{
			"#count": aws.String("count"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":start": {
				N: aws.String("0"),
			},
			":num": {
				N: aws.String("1"),
			},
		},
		ReturnValues: aws.String("ALL_NEW"),
	}
	output, err := svc.UpdateItem(input)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
		}, err
	}
	counter := &models.Counter{}
	err = dynamodbattribute.UnmarshalMap(output.Attributes, counter)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
		}, err
	}
	counterJSON, err := json.Marshal(counter)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
		}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(counterJSON),
	}, nil
}

func main() {
	lambda.Start(Handler)
}
