package lookup

import (
	"os"
	"path/filepath"
)

type OCR struct {
	fontFamilies map[string][]*FontSymbol
	threshold    float32
}

func NewOCR(threshold float32) *OCR {
	ocr := &OCR{
		fontFamilies: make(map[string][]*FontSymbol),
		threshold:    threshold,
	}

	return ocr
}

func (o *OCR) loadFont(fontPath string) error {
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
		family = make([]*FontSymbol, len(fontFamily))
		o.fontFamilies[familyName] = family
	}

	family = append(family, fontFamily...)
	return nil
}
