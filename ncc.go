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

func LookupAll(img image.Image, template image.Image, m float64) ([]GPoint, error) {
	imgBin := common.NewImageBinaryGrey(img)
	templateBin := common.NewImageBinaryGrey(template)

	var list []GPoint
	x1, y1 := 0, 0
	max := img.Bounds().Max
	x2, y2 := max.X-1, max.Y-1

	templateMax := template.Bounds().Max
	for x := x1; x <= x2-templateMax.X+1; x++ {
		for y := y1; y <= y2-templateMax.Y+1; y++ {
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
	m := common.NewImageMultiply(img.ZeroMeanImage(), xx, yy, template.ZeroMeanImage())
	return m.Sum()
}
