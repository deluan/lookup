package lookup

import (
	"image"
	"image/color"
)

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
// 2) Integral ^ 2 Image (Image Energy)
// 3) Zero Mean Image (image where each pixel subtracted with image mean value)
type ImageBinaryChannel struct {
	ChannelType   ChannelType
	ZeroMeanImage []float64
	integralImage *IntegralImage
	Width         int
	Height        int
}

// Standard deviation, no sqrt and no mean
func (c *ImageBinaryChannel) Dev2nRect(x1, y1, x2, y2 int) float64 {
	return c.integralImage.dev2nRect(x1, y1, x2, y2)
}

// Same as Dev2nRect, for the whole image
func (c *ImageBinaryChannel) Dev2n() float64 {
	return c.integralImage.dev2nRect(0, 0, c.Width-1, c.Height-1)
}

// Container for ImageBinaryChannels (one for each channel)
// It auto-detects if the image is RGB or GrayScale
type ImageBinary struct {
	Channels []*ImageBinaryChannel
	Width    int
	Height   int
	Size     int
}

func NewImageBinary(img image.Image) *ImageBinary {
	max := img.Bounds().Max
	ib := &ImageBinary{
		Width:  max.X,
		Height: max.Y,
		Size:   max.X * max.Y,
	}
	if _, ok := img.(*image.Gray); ok {
		c := newImageBinaryChannel(img, Gray)
		ib.Channels = append(ib.Channels, c)
	} else {
		ib.Channels = newImageBinaryChannels(img, Red, Green, Blue)
	}
	return ib
}

func newImageBinaryChannel(img image.Image, channelType ChannelType) *ImageBinaryChannel {
	max := img.Bounds().Max
	ibc := &ImageBinaryChannel{
		ChannelType: channelType,
		Width:       max.X,
		Height:      max.Y,
	}
	ibc.integralImage = NewIntegralImage(img)
	ibc.ZeroMeanImage = createZeroMeanImage(img, ibc.integralImage.Mean)

	return ibc
}

// Extract one or more color channels from image, and wraps them in ImageBinaryChannels
func newImageBinaryChannels(imgSrc image.Image, colorChannelTypes ...ChannelType) []*ImageBinaryChannel {
	channels := make([]*ImageBinaryChannel, 3)
	max := imgSrc.Bounds().Max
	w, h := max.X, max.Y
	for i, channelType := range colorChannelTypes {
		colorChannel := image.NewGray(image.Rectangle{Min: image.Point{0, 0}, Max: image.Point{w, h}})
		for x := 0; x < w; x++ {
			for y := 0; y < h; y++ {
				colorPixel := imgSrc.At(x, y).(color.NRGBA)
				var c uint8
				switch channelType {
				case Red:
					c = colorPixel.R
				case Green:
					c = colorPixel.G
				case Blue:
					c = colorPixel.B
				}
				grayPixel := color.Gray{Y: c}
				colorChannel.Set(x, y, grayPixel)
			}
		}
		channels[i] = newImageBinaryChannel(colorChannel, channelType)
	}
	return channels
}

// Zero Mean Image is an image where every pixel in the image is subtracted by
// mean value of the image. Mean value of the image is sum of all pixels values
// divided by number of pixels.
func createZeroMeanImage(img image.Image, mean float64) []float64 {
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
