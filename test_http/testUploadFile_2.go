package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

// Creates a new file upload http request with optional extra params
func newfileUploadRequest(uri string, params map[string]string, paramName, path string) (*http.Request, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(paramName, filepath.Base(path))
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", uri, body)
	contentType := writer.FormDataContentType()
	request.Header.Set("Content-Type", contentType)
	return request, err
}

func main() {
	target_url := "http://localhost:80/products"
	filename := "./golang.png"

	path, _ := os.Getwd()
	path += "/test.pdf"
	fmt.Println(path)

	extraParams := map[string]string{
		"name":     "XBox",
		"price":    "70000",
		"provider": "Microsoft",
		"rating":   "3.5",
		"status":   "sale",
	}

	request, err := newfileUploadRequest(target_url, extraParams, "image", filename)
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	} else {
		body := &bytes.Buffer{}
		_, err := body.ReadFrom(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		resp.Body.Close()
		fmt.Println(resp.StatusCode)
		fmt.Println(resp.Header)

		fmt.Println(body)
	}
}
