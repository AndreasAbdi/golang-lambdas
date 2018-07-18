package main

// this'll be a test on storing objects to s3.
import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type Request struct {
	ID    float64 `json:"id"`
	value string  `json:"value"`
}

type Response struct {
	Message string `json:"message"`
	Ok      bool   `json:"ok"`
}

func StoreHandler(request Request) (Response, error) {
	stored := false
	const storedwe = "fwef"

	if stored {
		return Response{
			Message: fmt.Sprintf("Process Request ID %f completed, stored to s3 bucket", request.ID),
			Ok:      true,
		}, nil
	} else {
		return Response{
			Message: fmt.Sprintf("Process Request ID %f not completed. Unable to store.", request.ID),
			Ok:      false,
		}, nil
	}
}

func main() {
	lambda.Start(StoreHandler)
}
