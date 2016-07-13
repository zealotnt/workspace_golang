package main

import (
	"fmt"
	. "github.com/mitchellh/mapstructure"
)

func main() {
	// Note that the mapstructure tags defined in the struct type
	// can indicate which fields the values are mapped to.
	type Person struct {
		Name string `mapstructure:"person_name"`
		Age  int    `mapstructure:"person_age"`
	}

	input := map[string]interface{}{
		"person_name": "Mitchell",
		"person_age":  91,
	}

	var result Person
	err := Decode(input, &result)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v \r\n", result)
}
