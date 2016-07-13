package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Previous env val: ", os.Getenv("S3_BUCKET_NAME"))

	os.Setenv("S3_BUCKET_NAME", "TestValue")

	fmt.Println("New env val: ", os.Getenv("S3_BUCKET_NAME"))
}
