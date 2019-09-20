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
			pixel := imgSrc.At(x, y)
			if _, ok := pixel.(color.NRGBA); ok {
				pixel = nrgbaToGray(pixel)
			} else {
				pixel = color.GrayModel.Convert(pixel)
			}
			grayImage.Set(x, y, pixel)
		}
	}
	return grayImage
}

func nrgbaToGray(pixel color.Color) color.Gray {
	p := pixel.(color.NRGBA)
	m := (float64(p.R) + float64(p.G) + float64(p.B)) / 3
	return color.Gray{Y: uint8(m)}
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
