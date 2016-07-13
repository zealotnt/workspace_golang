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

	for idx := range resp.Contents {
		fmt.Println(*resp.Contents[idx].Key)

		params := &s3.DeleteObjectsInput{
			Bucket: aws.String(os.Getenv("S3_BUCKET_NAME")), // Required
			Delete: &s3.Delete{ // Required
				Objects: []*s3.ObjectIdentifier{ // Required
					{ // Required
						Key:       aws.String(*resp.Contents[idx].Key), // Required
						// VersionId: aws.String("ObjectVersionId"),
					},
					// More values...
				},
				Quiet: aws.Bool(true),
			},
			// MFA:          aws.String("MFA"),
			// RequestPayer: aws.String("RequestPayer"),
		}

		_, err := svc.DeleteObjects(params)

		if err != nil {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
			return
		}
	}
}
