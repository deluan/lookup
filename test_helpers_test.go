package lookup

import (
	"image"
	"os"
)

func loadImageColor(path string) image.Image {
	imageFile, _ := os.Open(path)
	defer imageFile.Close()
	img, _, _ := image.Decode(imageFile)
	return img
}

func loadImageGray(path string) image.Image {
	img := loadImageColor(path)
	return EnsureGrayScale(img)
}

func newGrayImage(width, height int, pixels []uint8) image.Image {
	grayImage := image.NewGray(image.Rect(0, 0, width, height))
	for i, v := range pixels {
		grayImage.Pix[i] = v
	}
	return grayImage
}
