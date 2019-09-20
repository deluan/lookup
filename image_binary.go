package lookup

import (
	"image"
	"image/color"
)

type channelType int

const (
	gray channelType = iota
	red
	green
	blue
)

// Minimal Sum-Tables required for any image for NCC or FNCC algorithm.
//
// 1) Integral Image
// 2) Integral ^ 2 Image (Image Energy)
// 3) Zero Mean Image (image where each pixel subtracted with image mean value)
type imageBinaryChannel struct {
	channelType   channelType
	zeroMeanImage []float64
	integralImage *integralImage
	width         int
	height        int
}

// Standard deviation, no sqrt and no mean
func (c *imageBinaryChannel) dev2nRect(x1, y1, x2, y2 int) float64 {
	return c.integralImage.dev2nRect(x1, y1, x2, y2)
}

// Same as Dev2nRect, for the whole image
func (c *imageBinaryChannel) dev2n() float64 {
	return c.integralImage.dev2nRect(0, 0, c.width-1, c.height-1)
}

// Container for ImageBinaryChannels (one for each channel)
// It auto-detects if the image is RGB or GrayScale
type imageBinary struct {
	channels []*imageBinaryChannel
	width    int
	height   int
	size     int
}

func newImageBinary(img image.Image) *imageBinary {
	max := img.Bounds().Max
	ib := &imageBinary{
		width:  max.X,
		height: max.Y,
		size:   max.X * max.Y,
	}
	if _, ok := img.(*image.Gray); ok {
		c := newImageBinaryChannel(img, gray)
		ib.channels = append(ib.channels, c)
	} else {
		ib.channels = newImageBinaryChannels(img, red, green, blue)
	}
	return ib
}

func newImageBinaryChannel(img image.Image, channelType channelType) *imageBinaryChannel {
	max := img.Bounds().Max
	ibc := &imageBinaryChannel{
		channelType: channelType,
		width:       max.X,
		height:      max.Y,
	}
	ibc.integralImage = newIntegralImage(img)
	ibc.zeroMeanImage = createZeroMeanImage(img, ibc.integralImage.mean)

	return ibc
}

// Extract one or more color channels from image, and wraps them in ImageBinaryChannels
func newImageBinaryChannels(imgSrc image.Image, colorChannelTypes ...channelType) []*imageBinaryChannel {
	channels := make([]*imageBinaryChannel, 3)
	max := imgSrc.Bounds().Max
	w, h := max.X, max.Y
	for i, channelType := range colorChannelTypes {
		colorChannel := image.NewGray(image.Rectangle{Min: image.Point{0, 0}, Max: image.Point{w, h}})
		for x := 0; x < w; x++ {
			for y := 0; y < h; y++ {
				colorPixel := imgSrc.At(x, y).(color.NRGBA)
				var c uint8
				switch channelType {
				case red:
					c = colorPixel.R
				case green:
					c = colorPixel.G
				case blue:
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
