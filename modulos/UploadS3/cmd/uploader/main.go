package main

import (
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var (
	s3Client *s3.S3
	s3Bucket string
	wg       sync.WaitGroup
)

func init() {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-1"),
		Credentials: credentials.NewStaticCredentials(
			"your-access-key-id",
			"your-secret",
			"your-session-token",
		),
	})

	if err != nil {
		panic(err)
	}

	s3Client = s3.New(sess)
	s3Bucket = "goexpert-bucket-exemplo"
}

func main() {
	dir, err := os.Open("./tmp")
	if err != nil {
		panic(err)
	}
	defer dir.Close()

	uploadControl := make(chan struct{}, 100)
	defer close(uploadControl)

	errorFileUpload := make(chan string, 10)

	go func() {
		for {
			select {
			case fileName := <-errorFileUpload:
				uploadControl <- struct{}{}
				wg.Add(1)
				go UploadFile(fileName, uploadControl, errorFileUpload)
			}
		}
	}()
	for {
		files, err := dir.Readdir(1)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Printf("Error reading directory, %v\n", err)
			continue
		}
		wg.Add(1)
		uploadControl <- struct{}{}
		go UploadFile(files[0].Name(), uploadControl, errorFileUpload)
	}
	wg.Wait()
}

func UploadFile(filename string, uploadControl <-chan struct{}, errorFileUpload chan<- string) {
	defer wg.Done()
	completeFileName := fmt.Sprintf("./tmp/%s", filename)
	fmt.Printf("Uploading file %q to bucket %q\n", completeFileName, s3Bucket)
	file, err := os.Open(completeFileName)
	if err != nil {
		fmt.Printf("failed to open file %q, %v\n", completeFileName, err)
		<-uploadControl
		errorFileUpload <- filename
		return
	}
	defer file.Close()

	_, err = s3Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(s3Bucket),
		Key:    aws.String(filename),
		Body:   file,
	})

	if err != nil {
		fmt.Printf("failed to upload file, %v\n", err)
		<-uploadControl
		errorFileUpload <- filename
		return
	}
	fmt.Printf("file %q uploaded to bucket %q\n", filename, s3Bucket)
	<-uploadControl
}
