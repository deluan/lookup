package utils

import (
	"image"
	"image/color"
)

// Why this does not work? https://stackoverflow.com/a/42518487
func ConvertToAverageGrayScale(imgSrc image.Image) image.Image {
	bounds := imgSrc.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y
	grayScale := image.NewGray(image.Rectangle{Min: image.Point{0, 0}, Max: image.Point{w, h}})
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			pixel := imgSrc.At(x, y).(color.NRGBA)
			m := (float64(pixel.R) + float64(pixel.G) + float64(pixel.B)) / 3
			grayColor := color.Gray{Y: uint8(m)}
			grayScale.Set(x, y, grayColor)
		}
	}
	return grayScale
}
