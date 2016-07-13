package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func myS3Download(key string) ([]byte, error) {
	downloader := s3manager.NewDownloader(session.New(&aws.Config{
		Region: aws.String(os.Getenv("S3_REGION")),
	}))

	var aws_buff aws.WriteAtBuffer
	_, err := downloader.Download(&aws_buff,
		&s3.GetObjectInput{
			Bucket: aws.String((os.Getenv("S3_BUCKET_NAME"))),
			Key:    aws.String(key),
	})

	return aws_buff.Bytes(), err
}

func main() {
	byteForm, err := myS3Download("docs/myKey")
	if err != nil {
		panic(err)
	}
	fmt.Println("Byte form: \r\n", byteForm, len(byteForm), "bytes")

	textForm := string(byteForm[:len(byteForm)])
	fmt.Println("Text form: \r\n", textForm)
}
