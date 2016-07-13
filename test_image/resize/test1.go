package main

import (
	"github.com/nfnt/resize"
	// "image/jpeg"
	"image/png"
	"log"
	"os"
)

func main() {
	// open "test.jpg"
	file, err := os.Open("./image.png")
	if err != nil {
		log.Fatal(err)
	}

	// decode jpeg into image.Image
	img, err := png.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	// resize to width 1000 using Lanczos resampling
	// and preserve aspect ratio
	// m := resize.Resize(1000, 0, img, resize.Lanczos3)
	m := resize.Resize(1000, 0, img, resize.MitchellNetravali)

	out, err := os.Create("image_resized.png")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// write new image to file
	png.Encode(out, m)
}
