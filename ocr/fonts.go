package ocr

import "github.com/deluan/lookup/common"

type FontSymbol struct {
	Symbol string
	image  common.ImageBinaryGrey
	Width  int
	Height int
}

func NewFontSymbol(symbol string, image common.ImageBinaryGrey) *FontSymbol {
	fs := &FontSymbol{
		Symbol: symbol,
		image:  image,
		Width:  image.Width(),
		Height: image.Height(),
	}

	return fs
}
