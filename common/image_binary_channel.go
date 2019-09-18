package common

// Minimal Sum-Tables required for any image for NCC or FNCC algorithm.
//
// 1) Base Image Array
//
// 2) Integral Image
//
// 3) Integral ^ 2 Image (Image Energy)
//
// 4) Zero Mean Image (image where each pixel substracted with image mean value)

type ChannelType int

const (
	Gray = iota
	Red
	Green
	Blue
)

type ImageBinaryChannel struct {
	basicSArray
	ChannelType ChannelType
	gi          SArray
	integral    *IntegralImage
	integral2   *IntegralImage2
	zeroMean    *ZeroMeanImage
}

func NewImageBinaryChannel(channelType ChannelType, img SArray) *ImageBinaryChannel {
	ibc := &ImageBinaryChannel{
		gi:          img,
		integral:    NewIntegralImageFromBase(img),
		integral2:   NewIntegralImage2FromBase(img),
		ChannelType: channelType,
	}
	stepThrough(ibc)
	ibc.zeroMean = NewZeroMeanImageFromBase(ibc.integral)
	return ibc
}

func (ibc *ImageBinaryChannel) Step(x, y int) {
	ibc.integral.Step(x, y)
	ibc.integral2.Step(x, y)
}

// Standard deviation no sqrt and no mean
func (ibc *ImageBinaryChannel) Dev2n() float64 {
	return ibc.integral2.dev2n(ibc.integral)
}

func (ibc *ImageBinaryChannel) Dev2nRect(x1, y1, x2, y2 int) float64 {
	return ibc.integral2.dev2nRect(ibc.integral, x1, y1, x2, y2)
}

func (ibc *ImageBinaryChannel) Width() int {
	return ibc.gi.Width()
}

func (ibc *ImageBinaryChannel) Height() int {
	return ibc.gi.Height()
}

func (ibc *ImageBinaryChannel) Size() int {
	return ibc.gi.Size()
}

func (ibc *ImageBinaryChannel) ZeroMeanImage() *ZeroMeanImage {
	return ibc.zeroMean
}
