package lookup

import (
	"fmt"
	"github.com/deluan/lookup/common"
	"image"
	"math"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type GPoint struct {
	X, Y int
	G    float64
}

func LookupAllGrey(img image.Image, template image.Image, m float64) ([]GPoint, error) {
	imgBin := common.NewImageBinaryGrey(img)
	templateBin := common.NewImageBinaryGrey(template)
	return lookupAll(imgBin, templateBin, m)
}

func lookupAll(imgBin common.ImageBinary, templateBin common.ImageBinary, m float64) ([]GPoint, error) {
	var list []GPoint
	x1, y1 := 0, 0
	x2, y2 := imgBin.Width()-1, imgBin.Height()-1

	templateWidth := templateBin.Width()
	templateHeight := templateBin.Height()
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

func lookup(img common.ImageBinary, template common.ImageBinary, x int, y int, m float64) (*GPoint, error) {
	ci := img.Channels()
	ct := template.Channels()

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

func gamma(img *common.ImageBinaryChannel, template *common.ImageBinaryChannel, xx int, yy int) float64 {
	d := denominator(img, template, xx, yy)
	if d == 0 {
		return -1
	}

	n := numerator(img, template, xx, yy)
	return n / d
}

func denominator(img *common.ImageBinaryChannel, template *common.ImageBinaryChannel, xx int, yy int) float64 {
	di := img.Dev2nRect(xx, yy, xx+template.Width()-1, yy+template.Height()-1)
	dt := template.Dev2n()
	return math.Sqrt(di * dt)
}

func numerator(img *common.ImageBinaryChannel, template *common.ImageBinaryChannel, xx int, yy int) float64 {
	return multiplyAndSum(img.ZeroMeanImage(), xx, yy, template.ZeroMeanImage())
}

func multiplyAndSum(img common.SArray, xx, yy int, template common.SArray) float64 {
	cx := template.Width()
	cy := template.Height()
	var sum float64
	for x := 0; x < cx; x++ {
		for y := 0; y < cy; y++ {
			value := img.Get(xx+x, yy+y) * template.Get(x, y)
			sum += value
		}
	}
	return sum
}
