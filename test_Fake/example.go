package main

import (
	"fmt"
	"github.com/icrowley/fake"
)

func main() {
	name := fake.FirstName()
	fullname := fake.FullName()
	product := fake.Product()

	fmt.Println(name)
	fmt.Println(fullname)
	fmt.Println(product)
}
