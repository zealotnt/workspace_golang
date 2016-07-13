package main

import (
	"time"
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(time.Now().UTC().Unix())
	fmt.Println("products/<id>/origin." + strconv.FormatInt(time.Now().UTC().Unix(), 10))
}