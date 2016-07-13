package main

import (
	"bytes"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/s3"
	"net/http"
	"os"
)

func main() {

	// DO NOT PUT credentials in code for production usage!
	// see https://www.socketloop.com/tutorials/golang-setting-up-configure-aws-credentials-with-official-aws-sdk-go
	// on setting creds from environment or loading from file

	// the file location and load default profile
	creds := credentials.NewSharedCredentials("/<change this>/.aws/credentials", "default")

	_, err := creds.Get()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	aws.DefaultConfig.Region = "us-east-1" //<--- change this to yours

	config := &aws.Config{
		Region:           "",
		Endpoint:         "s3.amazonaws.com", // <-- forking important !
		S3ForcePathStyle: true,               // <-- without these lines. All will fail! fork you aws!
		Credentials:      creds,
		LogLevel:         0, // <-- feel free to crank it up
	}

	s3client := s3.New(config)

	bucketName := "<change>" // <-- change this to your bucket name

	fileToUpload := "<change>"

	file, err := os.Open(fileToUpload)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	fileInfo, _ := file.Stat()
	var size int64 = fileInfo.Size()

	buffer := make([]byte, size)

	// read file content to buffer
	file.Read(buffer)

	fileBytes := bytes.NewReader(buffer) // convert to io.ReadSeeker type

	fileType := http.DetectContentType(buffer)

	path := "/examplefolder/" + file.Name() // target file and location in S3

	params := &s3.PutObjectInput{
		Bucket:        aws.String(bucketName), // required
		Key:           aws.String(path),       // required
		ACL:           aws.String("public-read"),
		Body:          fileBytes,
		ContentLength: aws.Long(size),
		ContentType:   aws.String(fileType),
		Metadata: map[string]*string{
			"Key": aws.String("MetadataValue"), //required
		},
		// see more at http://godoc.org/github.com/aws/aws-sdk-go/service/s3#S3.PutObject
	}

	result, err := s3client.PutObject(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	fmt.Println(awsutil.StringValue(result))
}
