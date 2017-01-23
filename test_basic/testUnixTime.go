package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	fmt.Println(time.Now().UTC().Unix())
	fmt.Println("products/<id>/origin." + strconv.FormatInt(time.Now().UTC().Unix(), 10))
}
