package main

import (
	"bytes"
	"fmt"
	"image/png"
	"io/ioutil"

	"github.com/disintegration/imaging"
)

func GetThumbnail(file string, size int) []byte {
	srcImage, err := imaging.Open(file)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	dstImage := imaging.Fill(srcImage, size, size, imaging.Center, imaging.Lanczos)
	// dstImage := imaging.Fit(srcImage, size, size, imaging.Lanczos)

	dataBytes := new(bytes.Buffer)

	if err := png.Encode(dataBytes, dstImage); err != nil {
		fmt.Println(err)
		return nil
	}

	return dataBytes.Bytes()
}

func GetDetail(file string, size int) []byte {
	srcImage, err := imaging.Open(file)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	dstImage := imaging.Resize(srcImage, size, 0, imaging.Lanczos)

	dataBytes := new(bytes.Buffer)

	if err := png.Encode(dataBytes, dstImage); err != nil {
		fmt.Println(err)
		return nil
	}

	return dataBytes.Bytes()
}

func main() {
	path := string("./1.jpg")
	thumbnail := GetThumbnail(path, 320)
	detail := GetDetail(path, 550)
	ioutil.WriteFile("thumbnail.png", thumbnail, 0777)
	ioutil.WriteFile("detail.png", detail, 0777)
}
