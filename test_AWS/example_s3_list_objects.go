package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"os"
	"fmt"
)

func main() {
	svc := s3.New(session.New(&aws.Config{Region: aws.String(os.Getenv("S3_REGION"))}))

	params := &s3.ListObjectsInput{
		Bucket:       aws.String(os.Getenv("S3_BUCKET_NAME")), // Required
		// Delimiter:    aws.String("Delimiter"),
		// EncodingType: aws.String("EncodingType"),
		// Marker:       aws.String("Marker"),
		// MaxKeys:      aws.Int64(1),
		// Prefix:       aws.String("Prefix"),
	}
	resp, err := svc.ListObjects(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}
