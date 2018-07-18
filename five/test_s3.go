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

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func createFile(fileContent string) *os.File {
	data := []byte(fileContent)
	tempFile, err := ioutil.TempFile("", "")
	check(err)
	_, err = tempFile.Write(data)
	check(err)
	return tempFile
}

func main() {
	file := createFile("some_text")
	defer file.Close()

	activeSession := session.Must(session.NewSession())
	uploader := s3manager.NewUploader(activeSession)

	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(testBucketName),
		Key:    aws.String(testBucketKey),
		Body:   file,
	})

	check(err)

	fmt.Printf("File uploaded to, %s\n", result.Location)

}
