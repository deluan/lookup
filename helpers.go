package lookup

import (
	"image"
	"image/color"
)

// ensureGrayScale is a helper function to convert any image.Image to image.Gray, using a simple
// average of the color channels. Ignores luminosity.
func ensureGrayScale(imgSrc image.Image) image.Image {
	if _, ok := imgSrc.(*image.Gray); ok && (imgSrc.Bounds().Min == image.Point{}) {
		return imgSrc
	}
	min := imgSrc.Bounds().Min
	max := imgSrc.Bounds().Max
	mx, my := min.X, min.Y
	w, h := max.X-mx, max.Y-my
	grayImage := image.NewGray(image.Rectangle{Max: image.Point{w, h}})
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			pixel := imgSrc.At(mx+x, my+y)
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
