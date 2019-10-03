package lookup

import (
	"fmt"
	"image"
	"io/ioutil"
	"net/url"
	"os"
	"strings"
)

type fontSymbol struct {
	symbol string
	image  *imageBinary
	width  int
	height int
}

func newFontSymbol(symbol string, img image.Image) *fontSymbol {
	imgBin := newImageBinary(img)
	fs := &fontSymbol{
		symbol: symbol,
		image:  imgBin,
		width:  imgBin.width,
		height: imgBin.height,
	}

	return fs
}

func (f *fontSymbol) String() string { return f.symbol }

type fontSymbolLookup struct {
	fs   *fontSymbol
	x, y int
	g    float64
	size int
}

func newFontSymbolLookup(fs *fontSymbol, x, y int, g float64) *fontSymbolLookup {
	return &fontSymbolLookup{fs, x, y, g, fs.image.size}
}

func (l *fontSymbolLookup) cross(f *fontSymbolLookup) bool {
	r := image.Rect(l.x, l.y, l.x+l.fs.width, l.y+l.fs.height)
	r2 := image.Rect(f.x, f.y, f.x+f.fs.width, f.y+f.fs.height)

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

func loadFont(path string) ([]*fontSymbol, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	fonts := make([]*fontSymbol, len(files))
	for i, f := range files {
		if f.IsDir() || strings.HasPrefix(f.Name(), ".") {
			continue
		}
		fs, err := loadSymbol(path, f.Name())
		if err != nil {
			return nil, err
		}
		fonts[i] = fs
	}
	return fonts, nil
}

func loadSymbol(path string, fileName string) (*fontSymbol, error) {
	imageFile, err := os.Open(path + "/" + fileName)
	if err != nil {
		return nil, err
	}
	defer imageFile.Close()

	img, _, err := image.Decode(imageFile)
	if err != nil {
		return nil, err
	}

	symbolName, err := url.QueryUnescape(fileName)
	if err != nil {
		return nil, err
	}

	symbolName = strings.Replace(symbolName, "\u200b", "", -1) // Remove zero width spaces
	fs := newFontSymbol(
		strings.TrimSuffix(symbolName, ".png"),
		EnsureGrayScale(img),
	)
	return fs, nil
}
