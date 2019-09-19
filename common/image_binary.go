package common

import (
	"image"
	"image/color"
)

// Zero Mean Image is an image where every pixel in the image is subtracted by
// mean value of the image. Mean value of the image is sum of all pixels values
// divided by number of pixels.
func CreateZeroMeanImage(img image.Image, mean float64) []float64 {
	cx := img.Bounds().Max.X
	cy := img.Bounds().Max.Y
	zeroMeanImage := make([]float64, cx*cy)
	offset := 0
	for y := 0; y < cy; y++ {
		for x := 0; x < cx; x++ {
			pix := float64(img.At(x, y).(color.Gray).Y)
			zeroMeanImage[offset] = pix - mean
			offset++
		}
	}
	return zeroMeanImage
}

type ChannelType int

const (
	Gray ChannelType = iota
	Red
	Green
	Blue
)

// Minimal Sum-Tables required for any image for NCC or FNCC algorithm.
//
// 1) Integral Image
//
// 2) Integral ^ 2 Image (Image Energy)
//
// 3) Zero Mean Image (image where each pixel subtracted with image mean value)
type ImageBinaryChannel struct {
	ChannelType   ChannelType
	ZeroMeanImage []float64
	integralImage *IntegralImage
	Width         int
	Height        int
}

func NewImageBinaryChannel(img image.Image, channelType ChannelType) *ImageBinaryChannel {
	max := img.Bounds().Max
	ibc := &ImageBinaryChannel{
		ChannelType: channelType,
		Width:       max.X,
		Height:      max.Y,
	}
	ibc.integralImage = CreateIntegralImage(img)
	ibc.ZeroMeanImage = CreateZeroMeanImage(img, ibc.integralImage.Mean)

	return ibc
}

func (c *ImageBinaryChannel) Dev2nRect(x1, y1, x2, y2 int) float64 {
	return c.integralImage.dev2nRect(x1, y1, x2, y2)
}

func (c *ImageBinaryChannel) Dev2n() float64 {
	return c.integralImage.dev2n()
}

type ImageBinary struct {
	Channels []*ImageBinaryChannel
	Width    int
	Height   int
	Size     int
}

func NewImageBinary(img image.Image) *ImageBinary {
	// TODO if RGB then create one channel per color...
	max := img.Bounds().Max
	ibg := &ImageBinary{
		Width:  max.X,
		Height: max.Y,
		Size:   max.X * max.Y,
	}
	c := NewImageBinaryChannel(img, Gray)
	ibg.Channels = append(ibg.Channels, c)
	return ibg
}
