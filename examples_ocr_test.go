package lookup_test

import (
	"fmt"
	"image"
	"os"

	"github.com/deluan/lookup"
)

// Helper function to load an image from the filesystem
func loadImageFromFile(imgPath string) image.Image {
	imageFile, _ := os.Open(imgPath)
	defer imageFile.Close()
	img, _, _ := image.Decode(imageFile)
	return img
}

func Example_ocr() {
	// Create an OCR object with an accuracy of 0.7
	ocr := lookup.NewOCR(0.7)

	// Load a fontSet
	_ = ocr.LoadFont("testdata/font_1")

	// Load an image to recognize
	img := loadImageFromFile("testdata/test3.png")

	// Recognize text in image
	text, _ := ocr.Recognize(img)

	// Print the results
	fmt.Printf("Text found in image: %s\n", text)

	// Output:
	// Text found in image: 3662
	// 32€/€
}
