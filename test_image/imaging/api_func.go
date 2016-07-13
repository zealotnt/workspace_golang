package main

import (
	"bytes"
	"image/png"
	"io/ioutil"

	"github.com/disintegration/imaging"
)

func GetThumbnail(file string, size int) []byte {
	srcImage, err := imaging.Open(file)
	if err != nil {
		return nil
	}

	dstImage := imaging.Fit(srcImage, size, size, imaging.Lanczos)

	dataBytes := new(bytes.Buffer)

	if err := png.Encode(dataBytes, dstImage); err != nil {
		return nil
	}

	return dataBytes.Bytes()
}

func GetDetail(file string, size int) []byte {
	srcImage, err := imaging.Open(file)
	if err != nil {
		return nil
	}

	dstImage := imaging.Resize(srcImage, size, 0, imaging.Lanczos)

	dataBytes := new(bytes.Buffer)

	if err := png.Encode(dataBytes, dstImage); err != nil {
		return nil
	}

	return dataBytes.Bytes()
}

func main() {
	path := string("./golang_big.png")
	thumbnail := GetThumbnail(path, 320)
	detail := GetDetail(path, 550)
	ioutil.WriteFile("thumbnail.png", thumbnail, 0777)
	ioutil.WriteFile("detail.png", detail, 0777)
}
