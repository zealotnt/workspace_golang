package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"os"
)

// func main() {
// 	svc := s3.New(session.New(&aws.Config{Region: aws.String(os.Getenv("S3_REGION"))}))

// 	params := &s3.GetObjectAclInput{
// 		Bucket:       aws.String(os.Getenv("S3_BUCKET_NAME")), // Required
// 		Key:          aws.String("myKey"),  // Required
// 		RequestPayer: aws.String("RequestPayer"),
// 		// VersionId:    aws.String("ObjectVersionId"),
// 	}
// 	resp, err := svc.GetObjectAcl(params)

// 	if err != nil {
// 		// Print the error, cast err to awserr.Error to get the Code and
// 		// Message from an error.
// 		fmt.Println(err.Error())
// 		return
// 	}

// 	// Pretty-print the response data.
// 	fmt.Println(resp)
// }

func main() {
	svc := s3.New(session.New(&aws.Config{Region: aws.String(os.Getenv("S3_REGION"))}))

	params := &s3.PutObjectAclInput{
		Bucket: aws.String(os.Getenv("S3_BUCKET_NAME")),     // Required
		Key:    aws.String("products/39/origin.1461171263"), // Required
		// ACL:    aws.String("public-read"),
		AccessControlPolicy: &s3.AccessControlPolicy{
			Grants: []*s3.Grant{
				{ // Required
					Grantee: &s3.Grantee{
						Type: aws.String("Group"), // Required
						// DisplayName:  aws.String("DisplayName"),
						// EmailAddress: aws.String("EmailAddress"),
						// ID:           aws.String("ID"),
						URI: aws.String("http://acs.amazonaws.com/groups/global/AllUsers"),
					},
					Permission: aws.String("FULL_CONTROL"),
				},
				// More values...
			},
			// Owner: &s3.Owner{
			// 	DisplayName: aws.String("DisplayName"),
			// 	ID:          aws.String("ID"),
			// },
		},
		GrantFullControl: aws.String("Group"),
		// GrantRead:        aws.String("GrantRead"),
		// GrantReadACP:     aws.String("GrantReadACP"),
		// GrantWrite:       aws.String("GrantWrite"),
		// GrantWriteACP:    aws.String("GrantWriteACP"),
		// RequestPayer:     aws.String("RequestPayer"),
	}
	resp, err := svc.PutObjectAcl(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println("Some err:", err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}
