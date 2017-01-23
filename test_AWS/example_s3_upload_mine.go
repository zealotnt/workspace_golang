package main

import (
	"bytes"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"os"
)

func myS3Upload(key string, data []byte) (*s3manager.UploadOutput, error) {
	uploader := s3manager.NewUploader(session.New(&aws.Config{Region: aws.String(os.Getenv("S3_REGION"))}))

	return uploader.Upload(&s3manager.UploadInput{
		Body:   bytes.NewReader(data),
		Bucket: aws.String((os.Getenv("S3_BUCKET_NAME"))),
		Key:    aws.String(key),
	})
}

func main() {
	data := []byte{1, 2, 3, 4}
	// data := []byte("1234")
	output, err := myS3Upload("docs/myKey", data)
	if err != nil {
		panic(err)
	}
	fmt.Println(output.Location, output.VersionID, output.UploadID)
}
