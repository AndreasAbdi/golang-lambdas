package main

// this'll be a test on storing objects to s3.
import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

//Request is a structured request object for requests to s3 store lambda
type Request struct {
	ID         float64 `json:"id"`
	Value      string  `json:"value"`
	BucketName string  `json:"bucketName"`
	BucketKey  string  `json:"bucketKey"`
}

//Response is a structured response object for responses from s3 store lambda
type Response struct {
	Message string `json:"message"`
	Ok      bool   `json:"ok"`
}

//StoreHandler is lambda handler for store data to s3 requests.
func StoreHandler(request Request) (Response, error) {
	stored, err := writeToS3(request.Value, request.BucketName, request.BucketKey)

	if err != nil {
		return Response{
			Message: fmt.Sprintf("Process Request ID %f failed, error occurred", request.ID),
			Ok:      false,
		}, err
	}

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
