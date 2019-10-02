package lookup_test

import (
	"fmt"
	"image"
	"os"

	"github.com/deluan/lookup"
)

func ExampleNewLookup() {
	// Load full image
	imageFile, _ := os.Open("testdata/cyclopst1.png")
	defer imageFile.Close()
	img, _, _ := image.Decode(imageFile)

	// Create a lookup for that image
	l := lookup.NewLookup(img)

	// Load a template to search inside the image
	templateFile, _ := os.Open("testdata/cyclopst1.png")
	defer templateFile.Close()
	template, _, _ := image.Decode(templateFile)

	// Find all occurrences of the template in the image
	pp, _ := l.FindAll(template, 0.9)
	if len(pp) > 0 {
		fmt.Printf("Found %d matches:\n", len(pp))
		for _, p := range pp {
			fmt.Printf("- (%d, %d) with %f accuracy\n", p.X, p.Y, p.G)
		}
	} else {
		println("No matches found")
	}
}

func ExampleNewOCR() {
	// Create an OCR object with an accuracy of 0.7
	ocr := lookup.NewOCR(0.7)

	// Load a fontSet
	_ = ocr.LoadFont("testdata/font_1")

	// Load an image to recognize
	imageFile, _ := os.Open("testdata/test3.png")
	defer imageFile.Close()
	img, _, _ := image.Decode(imageFile)

	// Recognize text in image
	text, _ := ocr.Recognize(img)
	fmt.Printf("Text found in image: %s\n", text)
}
