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

// Creates a Lookout object. Lookout objects created by these function will do a gray scale search of the templates.
// Note that if the image or any template is not GrayScale, it will be converted automatically. This could cause some
// performance hits
func NewLookup(image image.Image) *Lookup {
	imgGray := ensureGrayScale(image)
	return &Lookup{
		imgBin: newImageBinary(imgGray),
	}
}

// Creates a Lookout object that works with all color channels of the image (RGB). Keep in mind that Lookout objects
// created with this function can only accept color templates as parameters to the FindAll method
func NewLookupColor(image image.Image) *Lookup {
	return &Lookup{
		imgBin: newImageBinary(image),
	}
}

//  Search for all occurrences of template inside image
func (l *Lookup) FindAll(template image.Image, threshold float64) ([]GPoint, error) {
	if len(l.imgBin.channels) == 1 {
		template = ensureGrayScale(template)
	}
	tb := newImageBinary(template)
	return lookupAll(l.imgBin, tb, threshold)
}
