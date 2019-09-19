package lookup

import (
	"image"
	"image/color"
)

// Converts a any image.Image to image.Gray, using a simple average of the color channels.
// Does not compute luminosity
// Ref: https://stackoverflow.com/a/42518487
func ConvertToAverageGrayScale(imgSrc image.Image) image.Image {
	max := imgSrc.Bounds().Max
	w, h := max.X, max.Y
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
