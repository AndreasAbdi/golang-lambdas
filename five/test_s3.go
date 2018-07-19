package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

const testBucketName = "aabdi.lambda.test"
const testBucketKey = "test.key.element"
const testRegion = "us-east-1"

func check(e error) {
	if e != nil {
		fmt.Printf("Error occured!!")
		panic(e)
	}
}

func createFile(fileContent string) string {
	data := []byte(fileContent)
	tempFile, err := ioutil.TempFile("", "")
	check(err)
	_, err = tempFile.Write(data)
	check(err)
	fmt.Printf("wrote to file named (%s): %s\n", tempFile.Name(), string(data))
	tempFile.Sync()
	return tempFile.Name()
}

func uploadFile(filename string) (*s3manager.UploadOutput, error) {
	file, err := os.Open(filename)
	check(err)
	awsConfig := aws.NewConfig().WithRegion(testRegion)
	activeSession := session.Must(session.NewSession(awsConfig))
	uploader := s3manager.NewUploader(activeSession)

	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(testBucketName),
		Key:    aws.String(testBucketKey),
		Body:   file,
	})

	return result, err
}

func main() {
	file := createFile("some_text")
	result, err := uploadFile(file)
	check(err)
	fmt.Printf("File uploaded to, %s\n", result.Location)

}
