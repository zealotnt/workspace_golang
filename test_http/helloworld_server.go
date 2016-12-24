package main

import (
	"fmt"
	"io"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

func main() {
	fmt.Println("Helloworld-Server serve at localhost:8000")

	http.HandleFunc("/", hello)
	http.ListenAndServe(":8000", nil)
}
