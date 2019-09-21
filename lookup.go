// Package lookup implements a nice, simple and fast library which helps you to
// lookup objects on a screen. It also includes OCR functionality. Using Lookup
// you can do OCR tricks like recognizing any information in your Robot application.
//
// The image search algorithm is based on the Normalized Cross Correlation algorithm.
package lookup

import "image"

type Lookup struct {
	imgBin *imageBinary
}

func NewLookup(image image.Image) *Lookup {
	return &Lookup{
		imgBin: newImageBinary(image),
	}
}

func NewLookupGrayScale(image image.Image) *Lookup {
	imgGray := ensureGrayScale(image)
	return NewLookup(imgGray)
}

//  Search for all occurrences of template inside image
func (l *Lookup) FindAll(template image.Image, threshold float64) ([]GPoint, error) {
	if len(l.imgBin.channels) == 1 {
		template = ensureGrayScale(template)
	}
	tb := newImageBinary(template)
	return lookupAll(l.imgBin, tb, threshold)
}
