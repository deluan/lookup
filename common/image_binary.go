package common

import "image"

type ImageBinary interface {
	Channels() []*ImageBinaryChannel
	Width() int
	Height() int
	Size() int
	Image() image.Image
}

type ImageBinaryGrey struct {
	SArray
	gi   *GrayImage
	gray *ImageBinaryChannel
}

func NewImageBinaryGrey(img image.Image) *ImageBinaryGrey {
	ibg := &ImageBinaryGrey{}
	ibg.gi = NewGreyImage(img)
	ibg.gray = NewImageBinaryChannel(Gray, ibg.gi)

	stepThrough(ibg)
	return ibg
}

func (ibg *ImageBinaryGrey) Step(x, y int) {
	ibg.gi.Step(x, y)
	ibg.gray.Step(x, y)
}

func (ibg *ImageBinaryGrey) Channels() []*ImageBinaryChannel {
	return []*ImageBinaryChannel{ibg.gray}
}

func (ibg *ImageBinaryGrey) Width() int {
	return ibg.gi.Width()
}

func (ibg *ImageBinaryGrey) Height() int {
	return ibg.gi.Height()
}

func (ibg *ImageBinaryGrey) Size() int {
	return ibg.gi.Size()
}

func (ibg *ImageBinaryGrey) Image() image.Image {
	return ibg.gi.buf
}
