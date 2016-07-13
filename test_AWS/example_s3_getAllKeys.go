package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"os"
)

func getAllKey() ([]string, error) {
	svc := s3.New(session.New(&aws.Config{Region: aws.String(os.Getenv("S3_REGION"))}))

	params := &s3.ListObjectsInput{
		Bucket: aws.String(os.Getenv("S3_BUCKET_NAME")), // Required
	}
	resp, err := svc.ListObjects(params)

	if err != nil {
		return nil, err
	}

	arr := []string{}
	for idx := range resp.Contents {
		arr = append(arr, *resp.Contents[idx].Key)
	}

	return arr, nil
}
func main() {
	AllKeys, _ := getAllKey()
	fmt.Println(AllKeys)
}
