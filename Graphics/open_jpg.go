package main

import (
	"fmt"
	"image/png"
	"os"
)

func main() {
	// Read image from file that already exists
	existingImageFile, err := os.Open("out.png")
	if err != nil {
		// Handle error
	}
	defer existingImageFile.Close()

	// Calling the generic image.Decode() will tell give us the data

	// We only need this because we already read from the file
	// We have to reset the file pointer back to beginning
	existingImageFile.Seek(0, 0)

	// Alternatively, since we know it is a png already
	// we can call png.Decode() directly
	loadedImage, err := png.Decode(existingImageFile)
	if err != nil {
		// Handle error
	}
  fmt.Println("imageImage:")
	fmt.Println((loadedImage))
}
