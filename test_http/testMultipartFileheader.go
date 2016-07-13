package main

import (
	"bytes"
	"fmt"
	"github.com/mholt/binding"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

type ProductForm struct {
	Name     *string               `json:"name"`
	Price    *int                  `json:"price"`
	Provider *string               `json:"provider"`
	Rating   *float32              `json:"rating"`
	Status   *string               `json:"status"`
	Image    *multipart.FileHeader `json:"image"`
}

func (form *ProductForm) FieldMap(req *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&form.Name: binding.Field{
			Form: "name",
		},
		&form.Price: binding.Field{
			Form: "price",
		},
		&form.Provider: binding.Field{
			Form: "provider",
		},
		&form.Rating: binding.Field{
			Form: "rating",
		},
		&form.Status: binding.Field{
			Form: "status",
		},
		&form.Image: binding.Field{
			Form: "image",
		},
	}
}

var image = new(multipart.FileHeader)

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

func (form ProductForm) isValidImage() bool {
	fh, err := form.Image.Open()
	if err != nil {
		return false
	}
	defer fh.Close()

	buff := make([]byte, 512)
	if _, err := fh.Read(buff); err != nil {
		return false
	}

	filetype := http.DetectContentType(buff)
	switch filetype {
	case "image/jpeg":
		fallthrough
	case "image/png":
		fallthrough
	case "image/gif":
		return true
	default:
		return false
	}

	return false
}

func (form *ProductForm) ImageData() []byte {
	if form.Image == nil {
		return nil
	}

	fh, err := form.Image.Open()
	if err != nil {
		panic(err)
	}
	defer fh.Close()

	dataBytes := bytes.Buffer{}

	dataBytes.ReadFrom(fh)

	return dataBytes.Bytes()
}

func main() {
	target_url := "http://localhost:80/products"
	filename := "./golang.png"
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

	// multipartReader, _ := request.MultipartReader()

	// form, _ := multipartReader.ReadForm(10 * 1024 * 1024)

	// request.MultipartForm = form

	product_form := new(ProductForm)
	if err := binding.Bind(request, product_form); err != nil {
		log.Fatal(err)
		return
	}
	// binding.Bind(request, product_form)

	fmt.Println(product_form.Image.Filename)
	fmt.Println(product_form.isValidImage())
	// fmt.Println(string(product_form.ImageData()))
}
