//  Search for all occurrences of template inside img, using the NCC algorithm.
package lookup

import "image"

func GrayScale(img image.Image, template image.Image, m float64) ([]GPoint, error) {
	imgGray := ensureGrayScale(img)
	templateGray := ensureGrayScale(template)
	return Color(imgGray, templateGray, m)
}

func Color(img image.Image, template image.Image, m float64) ([]GPoint, error) {
	imgBin := newImageBinary(img)
	templateBin := newImageBinary(template)
	return lookupAll(imgBin, templateBin, m)
}
