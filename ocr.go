package lookup

import (
	"image"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type OCR struct {
	fontFamilies map[string][]*fontSymbol
	threshold    float64
	allSymbols   []*fontSymbol
	numThreads   int
}

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

func (o *OCR) Recognize(img image.Image) (string, error) {
	bi := newImageBinary(ensureGrayScale(img))
	return o.recognize(bi, 0, 0, bi.width-1, bi.height-1)
}

func (o *OCR) recognize(bi *imageBinary, x1, y1, x2, y2 int) (string, error) {
	found, err := o.findAll(o.allSymbols, bi, x1, y1, x2, y2)
	if err != nil {
		return "", err
	}

	if len(found) == 0 {
		return "", nil
	}

	text := o.filterAndArrange(found)
	return text, nil
}

func (o *OCR) findAll(symbols []*fontSymbol, bi *imageBinary, x1, y1, x2, y2 int) ([]*fontSymbolLookup, error) {
	jc := startJob(o.numThreads)
	for _, fs := range symbols {
		jc.lookupSymbolParallel(bi, fs, o.threshold)
	}

	results, err := jc.collectResults()
	if err != nil {
		return nil, err
	}
	return results, nil
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
				// delete all[j]
				copy(all[j:], all[j+1:])
				all[len(all)-1] = nil
				all = all[:len(all)-1]
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
	for _, s := range all {
		maxCX := max(cx, s.fs.width)

		// if distance between end of previous symbol and beginning of the
		// current is larger then a char size, then it is a space
		if s.x-(x+cx) >= maxCX {
			str.WriteString(" ")
		}

		// if we drop back, then we have a end of line
		if s.x < x {
			str.WriteString("\n")
		}

		x = s.x + s.fs.width
		cx = s.fs.width
		str.WriteString(s.fs.symbol)
	}

	return str.String()
}
