package lookup

import "image"

// Lookup implements a image search algorithm based on Normalized Cross Correlation.
// For an overview of the algorithm, see http://www.fmwconcepts.com/imagemagick/similar/index.php
type Lookup struct {
	imgBin *imageBinary
}

// Creates a Lookup object. Lookup objects created by these function will do a gray
// scale search of the templates. Note that if the image or any template is not GrayScale,
// it will be converted automatically. This could cause some performance hits
func NewLookup(image image.Image) *Lookup {
	imgGray := EnsureGrayScale(image)
	return &Lookup{
		imgBin: newImageBinary(imgGray),
	}
}

// Creates a Lookup object that works with all color channels of the image (RGB).
// Keep in mind that Lookup objects created with this function can only accept color
// templates as parameters to the FindAll method
func NewLookupColor(image image.Image) *Lookup {
	return &Lookup{
		imgBin: newImageBinary(image),
	}
}

//  Search for all occurrences of template inside image
func (l *Lookup) FindAll(template image.Image, threshold float64) ([]GPoint, error) {
	if len(l.imgBin.channels) == 1 {
		template = EnsureGrayScale(template)
	}
	tb := newImageBinary(template)
	return lookupAll(l.imgBin, tb, threshold)
}
