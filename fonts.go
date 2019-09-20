package lookup

import (
	"image"
	"io/ioutil"
	"net/url"
	"os"
	"strings"
)

type FontSymbol struct {
	Symbol string
	image  *ImageBinary
	Width  int
	Height int
}

func NewFontSymbol(symbol string, img image.Image) *FontSymbol {
	imgBin := NewImageBinary(img)
	fs := &FontSymbol{
		Symbol: symbol,
		image:  imgBin,
		Width:  imgBin.Width,
		Height: imgBin.Height,
	}

	return fs
}

func (f *FontSymbol) String() string { return f.Symbol }

func loadFont(path string) ([]*FontSymbol, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	fonts := make([]*FontSymbol, len(files))
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

	symbolName, err := url.QueryUnescape(fileName)
	if err != nil {
		return nil, err
	}

	symbolName = strings.Replace(symbolName, "\u200b", "", -1) // Remove zero width spaces
	fs := NewFontSymbol(strings.TrimSuffix(symbolName, ".png"), ConvertToAverageGrayScale(img))
	return fs, nil
}
