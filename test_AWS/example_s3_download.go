package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	// "github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"log"
	"os"
)

func main() {
	file, err := os.Create("s3_download_file")
	if err != nil {
		log.Fatal("Failed to create file", err)
	}
	defer file.Close()

	downloader := s3manager.NewDownloader(session.New(&aws.Config{
		Region:      aws.String(os.Getenv("S3_REGION")),
		// Credentials: credentials.NewStaticCredentials(os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"), ""),
	}))

	numBytes, err := downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: aws.String((os.Getenv("S3_BUCKET_NAME"))),
			Key:    aws.String("myKey"),
		})
	if err != nil {
		fmt.Println("Failed to download file", err)
		return
	}

	fmt.Println("Downloaded file", file.Name(), numBytes, "bytes")
}
