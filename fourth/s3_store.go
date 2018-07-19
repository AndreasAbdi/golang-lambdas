package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

const testRegion = "us-east-1"

func check(e error) {
	if e != nil {
		fmt.Printf("Error occured!!")
		panic(e)
	}
}

func writeToS3(data string, bucketName string, bucketKey string) (bool, error) {
	fmt.Printf("Creating the file on s3\n")
	filename, err := createFile(data)
	if err != nil {
		return false, err
	}
	fmt.Printf("Uploading the file\n")
	err = uploadFile(filename, bucketName, bucketKey)
	if err != nil {
		return false, err
	}
	return true, nil
}

func createFile(fileContent string) (string, error) {
	data := []byte(fileContent)
	tempFile, err := ioutil.TempFile("", "")
	if err != nil {
		return "", err
	}
	_, err = tempFile.Write(data)
	if err != nil {
		return "", err
	}
	fmt.Printf("wrote to file named (%s): %s\n", tempFile.Name(), string(data))
	defer tempFile.Close()
	return tempFile.Name(), nil
}

func uploadFile(filename string, bucketName string, bucketKey string) error {
	fmt.Printf("Uploading to s3 bucket %s, at key %s\n", bucketName, bucketKey)
	fmt.Printf("Opening the file\n")
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	fmt.Printf("creating sessions\n")
	awsConfig := aws.NewConfig().WithRegion(testRegion)
	activeSession := session.Must(session.NewSession(awsConfig))
	uploader := s3manager.NewUploader(activeSession)

	fmt.Printf("uploading file to session\n")
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(bucketKey),
		Body:   file,
	})
	if err != nil {
		return err
	}
	return nil
}
