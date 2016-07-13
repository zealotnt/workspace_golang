package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"log"
	"os"
	"time"
)

func main() {
	svc := s3.New(session.New(&aws.Config{Region: aws.String(os.Getenv("S3_REGION"))}))
	req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(os.Getenv("S3_BUCKET_NAME")),
		Key:    aws.String("docs/myKey"),
	})
	urlStr, err := req.Presign(15 * time.Minute)
	if err != nil {
		log.Println("Failed to sign request", err)
	}
	log.Println("The URL is", urlStr)
}
