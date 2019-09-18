package common

// Zero Mean Image is an image where every pixel in the image is subtracted by
// mean value of the image. Mean value of the image is sum of all pixels values
// divided by number of pixels.

type ZeroMeanImage struct {
	basicSArray
	mean float64
}

func NewZeroMeanImageFromBase(base *IntegralImage) *ZeroMeanImage {
	zi := &ZeroMeanImage{}
	zi.initBase(base.base)
	zi.mean = base.Mean()
	stepThrough(zi)
	return zi
}

func (zi *ZeroMeanImage) Step(x, y int) {
	_ = zi.Set(x, y, zi.base.Get(x, y)-zi.mean)
}
