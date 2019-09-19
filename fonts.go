package lookup

import (
	"image"
)

type FontSymbol struct {
	Symbol string
	image  *ImageBinary
	Width  int
	Height int
}

func NewFontSymbol(symbol string, img image.Image) *FontSymbol {
	imgBin := NewImageBinary(img)
	fs := &FontSymbol{
		Symbol: symbol,
		image:  imgBin,
		Width:  imgBin.Width,
		Height: imgBin.Height,
	}

	return fs
}
