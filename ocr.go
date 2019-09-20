package lookup

import (
	"fmt"
	"image"
	"os"
	"path/filepath"
	"sort"
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
	bi := NewImageBinary(ConvertToAverageGrayScale(img))
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

	text := o.selectBestMatches(found)
	return text, nil
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
	size       int
}

func newFontSymbolLookup(fs *FontSymbol, x, y int, g float64) *FontSymbolLookup {
	return &FontSymbolLookup{fs, x, y, g, fs.image.Size}
}

func (l *FontSymbolLookup) cross(f *FontSymbolLookup) bool {
	r := image.Rect(l.x, l.y, l.x+l.fontSymbol.Width, l.y+l.fontSymbol.Height)
	r2 := image.Rect(f.x, f.y, f.x+f.fontSymbol.Width, f.y+f.fontSymbol.Height)

	return r.Intersect(r2) != image.Rectangle{}
}

func (l *FontSymbolLookup) yCross(f *FontSymbolLookup) bool {
	ly1 := l.y
	ly2 := l.y + l.fontSymbol.Height
	fy1 := f.y
	fy2 := f.y + f.fontSymbol.Height

	return (fy1 >= ly1 && fy1 <= ly2) || (fy2 >= ly1 && fy2 <= ly2)
}

func (l *FontSymbolLookup) bigger(other *FontSymbolLookup, maxSize2 int) bool {
	if abs(abs(l.size)-abs(other.size)) >= maxSize2 {
		return other.size < l.size
	}

	// better quality goes first
	diff := l.g - other.g
	if diff != 0 {
		return diff > 0
	}

	// bigger items goes first
	return other.size < l.size
}

func (l *FontSymbolLookup) follows(f *FontSymbolLookup) bool {
	r := 0
	if !l.yCross(f) {
		r = l.y - f.y
	}

	if r == 0 {
		r = l.x - f.x
	}

	if r == 0 {
		r = l.y - f.y
	}

	return r < 0
}

func (l *FontSymbolLookup) String() string {
	return fmt.Sprintf("'%s'(%d,%d,%d)[%f]", l.fontSymbol.Symbol, l.x, l.y, l.size, l.g)
}

func (o *OCR) findAll(symbols []*FontSymbol, bi *ImageBinary, x1, y1, x2, y2 int) ([]*FontSymbolLookup, error) {
	var found []*FontSymbolLookup

	for _, fs := range symbols {
		pp, err := lookupAll(bi, fs.image, o.threshold)
		if err != nil {
			return nil, err
		}
		for _, p := range pp {
			fsl := newFontSymbolLookup(fs, p.X, p.Y, p.G)
			found = append(found, fsl)
		}
	}

	return found, nil
}

func biggerFirst(list []*FontSymbolLookup) func(i, j int) bool {
	maxSize := 0
	for _, i := range list {
		maxSize = max(maxSize, i.fontSymbol.image.Size)
	}
	maxSize2 := maxSize / 2

	return func(i, j int) bool {
		return list[i].bigger(list[j], maxSize2)
	}
}

func (o *OCR) selectBestMatches(all []*FontSymbolLookup) string {
	var str strings.Builder

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
		return all[i].follows(all[j])
	})

	str.Grow(len(all) * 5)

	x := all[0].x
	cx := 0
	for _, s := range all {
		maxCX := max(cx, s.fontSymbol.Width)

		// if distance between end of previous symbol and beginning of the
		// current is larger then a char size, then it is a space
		if s.x-(x+cx) >= maxCX {
			str.WriteString(" ")
		}

		// if we drop back, then we have a end of line
		if s.x < x {
			str.WriteString("\n")
		}

		x = s.x + s.fontSymbol.Width
		cx = s.fontSymbol.Width
		str.WriteString(s.fontSymbol.Symbol)
	}

	return str.String()
}
