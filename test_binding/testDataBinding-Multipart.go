package main

import (
	"bytes"
	"fmt"
	"github.com/mholt/binding"
	"io"
	"log"
	"mime/multipart"
	"net/http"
)

// We expect a multipart/form-data upload with
// a file field named 'data'
type MultipartForm struct {
	Image *multipart.FileHeader `json:"data"`
}

func (f *MultipartForm) FieldMap(req *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&f.Image: "image",
	}
}

// Handlers are still clean and simple
func handler(resp http.ResponseWriter, req *http.Request) {
	multipartForm := new(MultipartForm)
	errs := binding.Bind(req, multipartForm)
	if errs.Handle(resp) {
		return
	}

	// To access the file data you need to Open the file
	// handler and read the bytes out.
	var fh io.ReadCloser
	var err error
	if fh, err = multipartForm.Image.Open(); err != nil {
		http.Error(resp,
			fmt.Sprint("Error opening Mime::Image %+v", err),
			http.StatusInternalServerError)
		return
	}
	defer fh.Close()
	dataBytes := bytes.Buffer{}
	var size int64
	if size, err = dataBytes.ReadFrom(fh); err != nil {
		http.Error(resp,
			fmt.Sprint("Error reading Mime::Image %+v", err),
			http.StatusInternalServerError)
		return
	}
	// Now you have the attachment in databytes.
	// Maximum size is default is 10MB.
	log.Printf("Read %v bytes with filename %s",
		size, multipartForm.Image.Filename)
}

func main() {
	http.HandleFunc("/upload", handler)
	http.ListenAndServe(":3000", nil)
}
