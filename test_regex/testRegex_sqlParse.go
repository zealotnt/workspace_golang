package main

import (
	"fmt"
	"regexp"
)

func main() {
	env := "mysql://root:1234@tcp(localhost:3306)/mytest"

	re, _ := regexp.Compile(`(\w*):\/\/(.+)`)
	result := re.FindStringSubmatch(env)
	fmt.Println(result[2])

}
