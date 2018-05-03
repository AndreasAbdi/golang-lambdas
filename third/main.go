package main

import (
	"errors"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

/*
https://github.com/aws/aws-lambda-go/tree/master/events

This area of the golang aws lambda is event types for the request objects.
Of note is that the request objects for auth'd requests have very interesting things.
apparently anon requests are different, need to check how the request is..

*/

var (
	// ErrNameNotProvided is thrown when a name is not provided
	ErrNameNotProvided = errors.New("no name was provided in the HTTP body")
)

func altHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	log.Printf("Processing Lambda request %n\n", request.RequestContext.RequestID)

	if len(request.Body) < 1 {
		return events.APIGatewayProxyResponse{}, ErrNameNotProvided
	}
	return events.APIGatewayProxyResponse{
		Body:       "Hello " + request.Body,
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(altHandler)
}
