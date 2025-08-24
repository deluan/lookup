package lookup

import (
	"fmt"
	"image"
	"io/ioutil"
	"math"
	"net/url"
	"os"
	"strings"
)

type FontSymbol struct {
	symbol  string
	image   *imageBinary
	width   int
	height  int
	advance int
}

// NewFontSymbolRune creates a new symbol for a rune. opts are optional (if set to nil).
func NewFontSymbolRune(symbol rune, img image.Image, opts *NewFontSymbolOptions) *FontSymbol {
	return NewFontSymbolOpts(string([]rune{symbol}), img, opts)
}

func NewFontSymbol(symbol string, img image.Image) *FontSymbol {
	return NewFontSymbolOpts(symbol, img, nil)
}

// NewFontSymbolOpts creates a new symbol for a rune. Use NewFontSymbol for using the default options.
func NewFontSymbolOpts(symbol string, img image.Image, opts *NewFontSymbolOptions) *FontSymbol {
	imgBin := newImageBinary(ensureGrayScale(img))
	advance := math.MaxInt
	if opts != nil {
		advance = opts.Advance
	}
	fs := &FontSymbol{
		symbol:  symbol,
		image:   imgBin,
		width:   imgBin.width,
		height:  imgBin.height,
		advance: advance,
	}

	return fs
}

func (f FontSymbol) Advance() int {
	if f.advance == math.MaxInt {
		return f.width
	}
	return f.advance
}

func (f *FontSymbol) String() string { return f.symbol }

type NewFontSymbolOptions struct {
	// The advance of the symbol, taken into account when recognizing texts./
	// This allows symbols to be closer/further away than the width of the symbol.
	// Is ignored when set to math.MaxInt
	Advance int
}

type fontSymbolLookup struct {
	fs   *FontSymbol
	x, y int
	g    float64
	size int
}

func newFontSymbolLookup(fs *FontSymbol, x, y int, g float64) *fontSymbolLookup {
	return &fontSymbolLookup{fs, x, y, g, fs.image.size}
}

func (l *fontSymbolLookup) cross(f *fontSymbolLookup) bool {
	r := image.Rect(l.x, l.y, l.x+l.fs.Advance(), l.y+l.fs.height)
	r2 := image.Rect(f.x, f.y, f.x+f.fs.Advance(), f.y+f.fs.height)

	return r.Intersect(r2) != image.Rectangle{}
}

func (l *fontSymbolLookup) yCross(f *fontSymbolLookup) bool {
	ly2 := l.y + l.fs.height
	fy2 := f.y + f.fs.height

	return (f.y >= l.y && f.y <= ly2) || (fy2 >= l.y && fy2 <= ly2)
}

func (l *fontSymbolLookup) biggerThan(other *fontSymbolLookup, maxSize2 int) bool {
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

func (l *fontSymbolLookup) comesAfter(f *fontSymbolLookup) bool {
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

func (l *fontSymbolLookup) String() string {
	return fmt.Sprintf("'%s'(%d,%d,%d)[%f]", l.fs.symbol, l.x, l.y, l.size, l.g)
}

func loadFont(path string) ([]*FontSymbol, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	fonts := make([]*FontSymbol, 0)
	for _, f := range files {
		if f.IsDir() || strings.HasPrefix(f.Name(), ".") {
			continue
		}
		fs, err := loadSymbol(path, f.Name())
		if err != nil {
			return nil, err
		}
		fonts = append(fonts, fs)
	}
	return fonts, nil
}

func loadSymbol(path string, fileName string) (*FontSymbol, error) {
	imageFile, err := os.Open(path + "/" + fileName)
	if err != nil {
		return nil, err
	}
	defer imageFile.Close()

	img, _, err := image.Decode(imageFile)
	if err != nil {
		return nil, err
	}

	nameWithoutExtension := strings.TrimSuffix(fileName, ".png")
	symbolName, err := url.QueryUnescape(nameWithoutExtension)
	if err != nil {
		return nil, err
	}

	symbolName = strings.Replace(symbolName, "\u200b", "", -1) // Remove zero width spaces
	fs := NewFontSymbol(
		symbolName,
		img,
	)
	return fs, nil
}
