package main

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"log"
	// "unsafe"
)

// The Vector type has unexported fields, which the package cannot access.
// We therefore write a BinaryMarshal/BinaryUnmarshal method pair to allow us
// to send and receive the type with the gob package. These interfaces are
// defined in the "encoding" package.
// We could equivalently use the locally defined GobEncode/GobDecoder
// interfaces.
type Vector struct {
	x, y, z int
}

func (v Vector) MarshalBinary() ([]byte, error) {
	// A simple encoding: plain text.
	b := new(bytes.Buffer)
	binary.Write(b, binary.LittleEndian, v.x)
	binary.Write(b, binary.LittleEndian, v.y)
	binary.Write(b, binary.LittleEndian, v.z)
	return b.Bytes(), nil
}

// UnmarshalBinary modifies the receiver so it must take a pointer receiver.
func (v *Vector) UnmarshalBinary(data []byte) error {
	var err error
	// A simple encoding: plain text.
	b := bytes.NewBuffer(data)
	binary.Read(b, binary.LittleEndian, &v.x)
	binary.Read(b, binary.LittleEndian, &v.y)
	binary.Read(b, binary.LittleEndian, &v.z)
	return err
}

// This example transmits a value that implements the custom encoding and decoding methods.
func main() {
	var network bytes.Buffer // Stand-in for the network.

	// Create an encoder and send a value.
	enc := gob.NewEncoder(&network)
	err := enc.Encode(Vector{3, 4, 5})
	if err != nil {
		log.Fatal("encode:", err)
	}
	spew.Println("Dump encoded data")
	fmt.Printf("%v\r\n", network)
	spew.Println(network.Len())

	// Create a decoder and receive a value.
	dec := gob.NewDecoder(&network)
	var v Vector
	err = dec.Decode(&v)
	if err != nil {
		log.Fatal("decode:", err)
	}

	fmt.Println("Dump decoded data")
	spew.Dump(v)
}
