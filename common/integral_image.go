package common

import "math"

//
// Sum Table. Integral sum.
//
// See http://en.wikipedia.org/wiki/Summed_area_table
//
type IntegralImage struct {
	basicSArray
}

func NewIntegralImageFromBase(base SArray) *IntegralImage {
	a := &IntegralImage{}
	a.initBase(base)
	stepThrough(a)
	return a
}

func (i *IntegralImage) Step(x, y int) {
	_ = i.Set(x, y, i.base.Get(x, y)+i.Get(x-1, y)+i.Get(x, y-1)-i.Get(x-1, y-1))
}

func (i *IntegralImage) SigmaRect(x1, y1, x2, y2 int) float64 {
	a := i.Get(x1-1, y1-1)
	b := i.Get(x2, y1-1)
	c := i.Get(x1-1, y2)
	d := i.Get(x2, y2)
	return a + d - b - c
}

func (i *IntegralImage) Sigma() float64 {
	return i.SigmaRect(0, 0, i.cx-1, i.cy-1)
}

func (i *IntegralImage) Mean() float64 {
	return i.Sigma() / float64(i.Size())
}

//
// Image Energy. Squared Image Function f^2(x,y).
//
// See http://en.wikipedia.org/wiki/Summed_area_table
//
type IntegralImage2 struct {
	IntegralImage
}

func NewIntegralImage2FromBase(base SArray) *IntegralImage2 {
	a := &IntegralImage2{}
	a.initBase(base)
	stepThrough(a)

	return a
}

func (i2 *IntegralImage2) Step(x, y int) {
	_ = i2.Set(x, y, math.Pow(i2.base.Get(x, y), 2)+i2.Get(x-1, y)+i2.Get(x, y-1)-i2.Get(x-1, y-1))
}

func pow(n float64) float64 {
	return n * n
}

func (i2 *IntegralImage2) dev2nRect(i *IntegralImage, x1, y1, x2, y2 int) float64 {
	sum := i.SigmaRect(x1, y1, x2, y2)
	size := (x2 - x1 + 1) * (y2 - y1 + 1)
	sum2 := i2.SigmaRect(x1, y1, x2, y2)
	result := sum2 - pow(sum)/float64(size)
	return result
}

func (i2 *IntegralImage2) dev2Rect(i *IntegralImage, x1, y1, x2, y2 int) float64 {
	size := (x2 - x1 + 1) * (y2 - y1 + 1)
	return i2.dev2nRect(i, x1, y1, x2, y2) / float64(size-1)
}

func (i2 *IntegralImage2) devRect(i *IntegralImage, x1, y1, x2, y2 int) float64 {
	return math.Sqrt(i2.dev2Rect(i, x1, y1, x2, y2))
}

// Standard deviation no sqrt and no mean
func (i2 *IntegralImage2) dev2n(i *IntegralImage) float64 {
	return i2.dev2nRect(i, 0, 0, i2.cx-1, i2.cy-1)
}

// Standard deviation no sqrt
func (i2 *IntegralImage2) dev2(i *IntegralImage) float64 {
	return i2.dev2Rect(i, 0, 0, i2.cx-1, i2.cy-1)
}

// Standard deviation
func (i2 *IntegralImage2) dev(i *IntegralImage) float64 {
	return i2.devRect(i, 0, 0, i2.cx-1, i2.cy-1)
}
