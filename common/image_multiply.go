package common

// Multiply matrix pixel by pixel
type ImageMultiply struct {
	basicSArray
	m      SArray
	xx, yy int
	sum    float64
}

func NewImageMultiply(image SArray, xx, yy int, template SArray) *ImageMultiply {
	a := &ImageMultiply{
		m:  image,
		xx: xx,
		yy: yy,
	}
	a.initBase(template)
	stepThrough(a)
	return a
}

func (im *ImageMultiply) Step(x, y int) {
	value := im.m.Get(im.xx+x, im.yy+y) * im.base.Get(x, y)
	_ = im.Set(x, y, value)
	im.sum += value
}

// Returns the sum of all multiplied pixels
func (im *ImageMultiply) Sum() float64 {
	return im.sum
}
