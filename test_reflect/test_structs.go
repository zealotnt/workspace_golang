package main

import (
	"github.com/fatih/structs"

	"encoding/json"
	"fmt"
	"net/http"
)

type Server struct {
	Name        string `json:"name,omitempty"`
	ID          int
	Enabled     bool
	users       []string // not exported
	http.Server          // embedded
}

func struct_method() {
	server := &Server{
		Name:    "gopher",
		ID:      123456,
		Enabled: true,
	}

	// Create a new struct type:
	s := structs.New(server)

	m := s.Map()               // Get a map[string]interface{}
	v := s.Values()            // Get a []interface{}
	fs := s.Fields()           // Get a []*Field
	ns := s.Names()            // Get a []string
	f := s.Field("Name")       // Get a *Field based on the given field name
	f, ok := s.FieldOk("Name") // Get a *Field based on the given field name
	n := s.Name()              // Get the struct name
	h := s.HasZero()           // Check if any field is initialized
	z := s.IsZero()            // Check if all fields are initialized

	tagValue := f.Tag("json") // Get the Field's tag value for tag name "json", tag value => "name,omitempty"

	fmt.Println(m)
	fmt.Println(v)
	fmt.Println(fs)
	fmt.Println(ns)
	fmt.Println(f)
	fmt.Println(ok)
	fmt.Println(n)
	fmt.Println(h)
	fmt.Println(z)
	fmt.Println(tagValue)

	m["Server"] = ""
	json_out, err := json.Marshal(m)
	fmt.Printf("\r\n\tJsonOutput:\r\n%s\r\n", json_out)
	fmt.Printf("\r\n\tJsonOutput:\r\n%s\r\n", err)
}

func test_basic() {
	server := &Server{
		Name:    "gopher",
		ID:      123456,
		Enabled: true,
	}

	// Convert a struct to a map[string]interface{}
	// => {"Name":"gopher", "ID":123456, "Enabled":true}
	m := structs.Map(server)
	fmt.Println(m)

	// Convert the values of a struct to a []interface{}
	// => ["gopher", 123456, true]
	v := structs.Values(server)
	fmt.Println(v)

	// Convert the names of a struct to a []string
	// (see "Names methods" for more info about fields)
	ns := structs.Names(server)
	fmt.Println(ns)

	// Convert the values of a struct to a []*Field
	// (see "Field methods" for more info about fields)
	f := structs.Fields(server)
	fmt.Println(f)

	// Return the struct name => "Server"
	n := structs.Name(server)
	fmt.Println(n)

	// Check if any field of a struct is initialized or not.
	h := structs.HasZero(server)
	fmt.Println(h)

	// Check if all fields of a struct is initialized or not.
	z := structs.IsZero(server)
	fmt.Println(z)

	// Check if server is a struct or a pointer to struct
	i := structs.IsStruct(server)
	fmt.Println(i)
}

func main() {
	// test_basic()
	struct_method()
}
