package lookup

import (
	"image"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// OCR implements a simple OCR based on the Lookup functions. It allows multiple fontsets,
// just call LoadFont for each fontset.
//
// If you need to encode special symbols use UNICODE in the file name. For example if you
// need to have '\' character (which is prohibited in the path and file name) specify
// %2F.png as a image symbol name.
//
// Sometimes you need to specify two different image for one symbol (if image / font symbol vary
// too much). To do so add unicode ZERO WIDTH SPACE symbol (%E2%80%8B) to the filename.
// Ex: %2F%E2%80%8B.png will produce '/' symbol as well.
type OCR struct {
	fontFamilies map[string][]*fontSymbol
	threshold    float64
	allSymbols   []*fontSymbol
	numThreads   int
}

// NewOCR creates a new OCR instance, that will use the given threshold. You can optionally
// parallelize the processing by specifying the number of threads to use. The optimal number
// varies and depends on your use case (size of fontset x size of image). Default is use
// only one thread
func NewOCR(threshold float64, numThreads ...int) *OCR {
	ocr := &OCR{
		fontFamilies: make(map[string][]*fontSymbol),
		threshold:    threshold,
		numThreads:   1,
	}

	if len(numThreads) > 0 {
		ocr.numThreads = numThreads[0]
	}

	return ocr
}

// LoadFont loads a specific fontset from the given folder. Fonts are simple image files
// containing a PNG/JPEG of the font, and named after the "letter" represented by the image.
//
// This can be called multiple times, with different folders, to load different fontsets.
func (o *OCR) LoadFont(fontPath string) error {
	if _, err := os.Stat(fontPath); os.IsNotExist(err) {
		return err
	}

	fontFamily, err := loadFont(fontPath)
	if err != nil {
		return err
	}

	familyName := filepath.Base(fontPath)
	family, ok := o.fontFamilies[familyName]
	if !ok {
		family = make([]*fontSymbol, 0, len(fontFamily))
	}

	family = append(family, fontFamily...)
	o.fontFamilies[familyName] = family

	o.updateAllSymbols()
	return nil
}

func (o *OCR) updateAllSymbols() {
	total := 0
	o.allSymbols = nil
	for _, family := range o.fontFamilies {
		total += len(family)
		o.allSymbols = append(o.allSymbols, family...)
	}
}

// Recognize the text in the image using the fontsets previously loaded. If a SubImage
// is received, the search will be limited by the boundaries of the SubImage
func (o *OCR) Recognize(img image.Image) (string, error) {
	bi := newImageBinary(ensureGrayScale(img))
	return o.recognize(bi, image.Rect(0, 0, bi.width-1, bi.height-1))
}

func (o *OCR) recognize(bi *imageBinary, rect image.Rectangle) (string, error) {
	found, err := findAllInParallel(o.numThreads, o.allSymbols, bi, o.threshold, rect)
	if err != nil {
		return "", err
	}

	if len(found) == 0 {
		return "", nil
	}

	text := o.filterAndArrange(found)
	return text, nil
}

func biggerFirst(list []*fontSymbolLookup) func(i, j int) bool {
	maxSize := 0
	for _, i := range list {
		maxSize = max(maxSize, i.fs.image.size)
	}
	maxSize2 := maxSize / 2

	return func(i, j int) bool {
		return list[i].biggerThan(list[j], maxSize2)
	}
}

func (o *OCR) filterAndArrange(all []*fontSymbolLookup) string {
	// big images eat small ones
	sort.Slice(all, biggerFirst(all))
	for k, kk := range all {
		for j := k + 1; j < len(all); j++ {
			jj := all[j]
			if kk.cross(jj) {
				all = deleteSymbol(all, j)
				j--
			}
		}
	}

	// sort top/bottom/left/right
	sort.Slice(all, func(i, j int) bool {
		return all[i].comesAfter(all[j])
	})

	var str strings.Builder
	x := all[0].x
	cx := 0
	for i, s := range all {
		maxCX := max(cx, s.fs.width)

		// if distance between end of previous symbol and beginning of the
		// current is larger then a char size, then it is a space
		// This should not be applied in the beginning (i == 0) as it would put a white space for
		// any s.x > maxCX will have a (useless) whitespace in front
		if s.x-x >= maxCX && i != 0 {
			str.WriteString(" ")
		}

		// if we drop back, then we have an end of line
		if s.x < x {
			str.WriteString("\n")
		}

		x = s.x + s.fs.width
		cx = s.fs.width
		str.WriteString(s.fs.symbol)
	}

	return str.String()
}

func deleteSymbol(all []*fontSymbolLookup, i int) []*fontSymbolLookup {
	copy(all[i:], all[i+1:])
	all[len(all)-1] = nil
	return all[:len(all)-1]
}
