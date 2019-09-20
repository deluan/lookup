package lookup

import (
	"image"
	"os"
	"path/filepath"
	"strings"
)

type OCR struct {
	fontFamilies map[string][]*FontSymbol
	threshold    float64
	totalSymbols int
}

func NewOCR(threshold float64) *OCR {
	ocr := &OCR{
		fontFamilies: make(map[string][]*FontSymbol),
		threshold:    threshold,
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
		family = make([]*FontSymbol, 0, len(fontFamily))
	}

	family = append(family, fontFamily...)
	o.fontFamilies[familyName] = family

	o.updateTotalSymbols()
	return nil
}

func (o *OCR) updateTotalSymbols() {
	total := 0
	for _, family := range o.fontFamilies {
		total += len(family)
	}
	o.totalSymbols = total
}

func (o *OCR) Recognize(img image.Image) (string, error) {
	bi := NewImageBinary(img)
	return o.recognize(bi, 0, 0, bi.Width-1, bi.Height-1)
}

func (o *OCR) recognize(bi *ImageBinary, x1, y1, x2, y2 int) (string, error) {
	symbols := o.getAllSymbols()
	found, err := o.findAll(symbols, bi, x1, y1, x2, y2)
	if err != nil {
		return "", err
	}

	if len(found) == 0 {
		return "", nil
	}

	var b strings.Builder
	b.Grow(len(found) * 5)
	for _, fs := range found {
		b.WriteString(fs.fontSymbol.Symbol)
	}
	return b.String(), nil
}

func (o *OCR) getAllSymbols() []*FontSymbol {
	symbols := make([]*FontSymbol, 0, o.totalSymbols)
	for _, family := range o.fontFamilies {
		symbols = append(symbols, family...)
	}
	return symbols
}

type FontSymbolLookup struct {
	fontSymbol *FontSymbol
	x, y       int
	g          float64
}

func (o *OCR) findAll(symbols []*FontSymbol, bi *ImageBinary, x1, y1, x2, y2 int) ([]*FontSymbolLookup, error) {
	var found []*FontSymbolLookup

	for _, fs := range symbols {
		pp, err := lookupAll(bi, fs.image, o.threshold)
		if err != nil {
			return nil, err
		}
		for _, p := range pp {
			fsl := &FontSymbolLookup{fs, p.X, p.Y, p.G}
			found = append(found, fsl)
		}
	}

	return found, nil
}
