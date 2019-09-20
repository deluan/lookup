package lookup

import (
	"image"
)

//
// Summed-area table
//
// See http://en.wikipedia.org/wiki/Summed_area_table
//
type integralImage struct {
	// Sum Table
	pix []float64
	// Image Energy. Squared Image Function f^2(x,y).
	pix2   []float64
	mean   float64
	width  int
	height int
}

func newIntegralImage(original image.Image) *integralImage {
	integral := createIntegralImage(original)
	integral.mean = integral.sigma(integral.pix, 0, 0, integral.width-1, integral.height-1) / float64(len(integral.pix))
	return integral
}

func (i *integralImage) get(pix []float64, x, y int) float64 {
	if x < 0 || y < 0 {
		return 0
	}
	idx := (y * i.width) + x
	return pix[idx]
}

func (i *integralImage) sigma(pixArray []float64, x1, y1, x2, y2 int) float64 {
	a := i.get(pixArray, x1-1, y1-1)
	b := i.get(pixArray, x2, y1-1)
	c := i.get(pixArray, x1-1, y2)
	d := i.get(pixArray, x2, y2)
	return a + d - b - c
}

func (i *integralImage) dev2nRect(x1, y1, x2, y2 int) float64 {
	sum := i.sigma(i.pix, x1, y1, x2, y2)
	size := (x2 - x1 + 1) * (y2 - y1 + 1)
	sum2 := i.sigma(i.pix2, x1, y1, x2, y2)
	result := sum2 - (sum*sum)/float64(size)
	return result
}

func createIntegralImage(original image.Image) *integralImage {
	max := original.Bounds().Max
	cx := max.X
	cy := max.Y
	integral := &integralImage{
		width:  cx,
		height: cy,
	}
	pix := make([]float64, cx*cy)
	pix2 := make([]float64, cx*cy)
	offset := 0
	originalGray := original.(*image.Gray).Pix
	for y := 0; y < cy; y++ {
		for x := 0; x < cx; x++ {
			a := float64(originalGray[offset])
			b := integral.get(pix, x-1, y)
			c := integral.get(pix, x, y-1)
			d := integral.get(pix, x-1, y-1)
			a2 := a * a
			b2 := integral.get(pix2, x-1, y)
			c2 := integral.get(pix2, x, y-1)
			d2 := integral.get(pix2, x-1, y-1)
			pix[offset] = a + b + c - d
			pix2[offset] = a2 + b2 + c2 - d2
			offset++
		}
	}
	integral.pix = pix
	integral.pix2 = pix2
	return integral
}
