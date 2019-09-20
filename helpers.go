package lookup

import (
	"image"
	"image/color"
)

// Converts a any image.Image to image.Gray, using a simple average of the color channels.
// Ignores luminosity
// Ref: https://stackoverflow.com/a/42518487
func ensureGrayScale(imgSrc image.Image) image.Image {
	if _, ok := imgSrc.(*image.Gray); ok {
		return imgSrc
	}
	max := imgSrc.Bounds().Max
	w, h := max.X, max.Y
	grayImage := image.NewGray(image.Rectangle{Min: image.Point{0, 0}, Max: image.Point{w, h}})
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			pixel := imgSrc.At(x, y).(color.NRGBA)
			m := (float64(pixel.R) + float64(pixel.G) + float64(pixel.B)) / 3
			grayColor := color.Gray{Y: uint8(m)}
			grayImage.Set(x, y, grayColor)
		}
	}
	return grayImage
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return x * -1
}
