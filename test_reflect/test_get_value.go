package main

import "fmt"
import "reflect"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	fmt.Println(GetField(&v, "Z"))
}

func GetField(v *Vertex, field string) int {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)
	if !f.IsValid() {
		fmt.Println(field + " is not part of struct")
		return 0
	}
	return int(f.Int())
}
