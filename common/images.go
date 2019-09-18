package common

import (
	"image"
	"image/color"
)

type GrayImage struct {
	basicSArray
	buf image.Image
}

func NewGreyImage(img image.Image) *GrayImage {
	gi := &GrayImage{buf: img}
	gi.cx = img.Bounds().Max.X
	gi.cy = img.Bounds().Max.Y
	gi.array = make([]float64, gi.cx*gi.cy)
	stepThrough(gi)
	return gi
}

func (gi *GrayImage) Step(x, y int) {
	pixel := gi.buf.At(x, y).(color.NRGBA)
	gray := int(pixel.R) + int(pixel.G) + int(pixel.B)
	_ = gi.Set(x, y, float64(gray/3))
}
