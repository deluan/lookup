package lookup

import (
	"fmt"
	"image"
	"math"
)

type GPoint struct {
	X, Y int
	G    float64
}

//  http://www.fmwconcepts.com/imagemagick/similar/index.php
//
//  1) mean && stddev
//  2) image1(x,y) - mean1 && image2(x,y) - mean2
//  3) [3] = (image1(x,y) - mean)(x,y)//  (image2(x,y) - mean)(x,y)
//  4) [4] = mean([3])
//  5) [4] / (stddev1//  stddev2)
//
//  Normalized Cross Correlation algorithm
func LookupAll(img image.Image, template image.Image, m float64) ([]GPoint, error) {
	imgBin := NewImageBinary(img)
	templateBin := NewImageBinary(template)
	return lookupAll(imgBin, templateBin, m)
}

func lookupAll(imgBin *ImageBinary, templateBin *ImageBinary, m float64) ([]GPoint, error) {
	var list []GPoint
	x1, y1 := 0, 0
	x2, y2 := imgBin.Width-1, imgBin.Height-1

	templateWidth := templateBin.Width
	templateHeight := templateBin.Height
	for x := x1; x <= x2-templateWidth+1; x++ {
		for y := y1; y <= y2-templateHeight+1; y++ {
			g, err := lookup(imgBin, templateBin, x, y, m)
			if err != nil {
				return nil, err
			}
			if g != nil {
				list = append(list, *g)
			}
		}
	}
	return list, nil
}

func lookup(img *ImageBinary, template *ImageBinary, x int, y int, m float64) (*GPoint, error) {
	ci := img.Channels
	ct := template.Channels

	ii := min(len(ci), len(ct))
	g := math.MaxFloat64

	for i := 0; i < ii; i++ {
		cct := ct[i]
		cci := ci[i]
		if cct.ChannelType != cci.ChannelType {
			return nil, fmt.Errorf("incompatible channels %d <> %d", cct.ChannelType, cci.ChannelType)
		}
		gg := gamma(cci, cct, x, y)
		if gg < m {
			return nil, nil
		}
		g = math.Min(g, gg)
	}
	return &GPoint{X: x, Y: y, G: g}, nil
}

func gamma(img *ImageBinaryChannel, template *ImageBinaryChannel, xx int, yy int) float64 {
	d := denominator(img, template, xx, yy)
	if d == 0 {
		return -1
	}

	n := numerator(img, template, xx, yy)
	return n / d
}

func denominator(img *ImageBinaryChannel, template *ImageBinaryChannel, xx int, yy int) float64 {
	di := img.Dev2nRect(xx, yy, xx+template.Width-1, yy+template.Height-1)
	dt := template.Dev2n()
	return math.Sqrt(di * dt)
}

func numerator(img *ImageBinaryChannel, template *ImageBinaryChannel, offsetX int, offsetY int) float64 {
	imgWidth := img.Width
	imgArray := img.ZeroMeanImage
	templateWidth := template.Width
	templateHeight := template.Height
	templateArray := template.ZeroMeanImage
	var sum float64
	for x := 0; x < templateWidth; x++ {
		for y := 0; y < templateHeight; y++ {
			value := imgArray[(offsetY+y)*imgWidth+offsetX+x] * templateArray[y*templateWidth+x]
			sum += value
		}
	}
	return sum
}
