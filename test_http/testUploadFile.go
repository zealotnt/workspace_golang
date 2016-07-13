// Reference:
// https://matt.aimonetti.net/posts/2013/07/01/golang-multipart-file-upload-example/
// https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/04.5.html
// https://golang.org/pkg/mime/multipart
// http://stackoverflow.com/questions/24493116/how-to-send-a-post-request-in-golang

package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
)

func postFile(filename string, targetUrl string) error {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)
	fileWriter, _ := bodyWriter.CreateFormFile("image", filename)
	fh, _ := os.Open(filename)
	io.Copy(fileWriter, fh)
	bodyWriter.WriteField("name", "xbox")
	bodyWriter.WriteField("price", "70000")
	bodyWriter.WriteField("provider", "Microsoft")
	bodyWriter.WriteField("rating", "3.5")
	bodyWriter.WriteField("status", "sale")
	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	fmt.Println(contentType)
	// fmt.Println(bodyBuf)

	resp, err := http.Post(targetUrl, contentType, bodyBuf)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(resp.Status)
	fmt.Println(string(resp_body))
	return nil
}

func myPost(filename string, targetUrl string) {
	response := httptest.NewRecorder()
	fmt.Println(response)

	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)
	fileWriter, _ := bodyWriter.CreateFormFile("image", "golang.png")
	fh, _ := os.Open("./golang.png")
	io.Copy(fileWriter, fh)
	bodyWriter.WriteField("name", "xbox")
	bodyWriter.WriteField("price", "70000")
	bodyWriter.WriteField("provider", "Microsoft")
	bodyWriter.WriteField("rating", "3.5")
	bodyWriter.WriteField("status", "sale")
	bodyWriter.Close()

	request, _ := http.NewRequest("POST", "http://localhost:80/products", bodyBuf)

	request.RemoteAddr = "http://localhost:80/"
	contentType := bodyWriter.FormDataContentType()
	request.Header.Set("Content-Type", contentType)

	client := http.Client{}
	resp2, err := client.Do(request)
	resp_body, err := ioutil.ReadAll(resp2.Body)
	defer resp2.Body.Close()

	fmt.Println(resp2)
	fmt.Println(resp2.Status)
	fmt.Println(string(resp_body))
	fmt.Println(err)
}

// sample usage
func main() {
	target_url := "http://localhost:80/products"
	filename := "./golang.go"
	myPost(filename, target_url)
	// postFile(filename, target_url)
}
