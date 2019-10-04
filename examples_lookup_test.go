package lookup_test

import (
	"fmt"
	"image"
	"os"

	"github.com/deluan/lookup"
)

// Helper function to load an image from the filesystem
func loadImage(imgPath string) image.Image {
	imageFile, _ := os.Open(imgPath)
	defer imageFile.Close()
	img, _, _ := image.Decode(imageFile)
	return img
}

func Example_lookup() {
	// Load full image
	img := loadImage("testdata/cyclopst1.png")

	// Create a lookup for that image
	l := lookup.NewLookup(img)

	// Load a template to search inside the image
	template := loadImage("testdata/cyclopst3.png")

	// Find all occurrences of the template in the image
	pp, _ := l.FindAll(template, 0.9)

	// Print the results
	if len(pp) > 0 {
		fmt.Printf("Found %d matches:\n", len(pp))
		for _, p := range pp {
			fmt.Printf("- (%d, %d) with %f accuracy\n", p.X, p.Y, p.G)
		}
	} else {
		println("No matches found")
	}

	// Output:
	// Found 1 matches:
	// - (21, 7) with 0.997942 accuracy
}
