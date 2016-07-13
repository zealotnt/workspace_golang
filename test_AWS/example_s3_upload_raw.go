package main

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func main() {
	file, err := os.Open("/home/zealot/Pictures/Wallpapers/4.jpg")
	if err != nil {
		log.Fatal("Failed to open file", err)
	}

	uploader := s3manager.NewUploader(session.New(&aws.Config{Region: aws.String(os.Getenv("S3_REGION"))}))
	result, err := uploader.Upload(&s3manager.UploadInput{
		Body:   file,
		Bucket: aws.String((os.Getenv("S3_BUCKET_NAME"))),
		Key:    aws.String("products/26/origin.1461404067.png"),
	})
	if err != nil {
		log.Fatalln("Failed to upload", err)
	}

	log.Println("Successfully uploaded to", result.Location)
}
