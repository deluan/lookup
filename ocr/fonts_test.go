package ocr

import (
	"image"
	_ "image/png"
	"os"
	"testing"

	"github.com/deluan/lookup/common"
	. "github.com/smartystreets/goconvey/convey"
)

func TestFontSymbol(t *testing.T) {
	Convey("Given an image and a template to look for", t, func() {
		img := loadImage("0.png")
		imgBin := common.NewImageBinaryGrey(img)
		fs := NewFontSymbol("0", imgBin)
		So(fs.Width, ShouldEqual, img.Bounds().Max.X)
		So(fs.Height, ShouldEqual, img.Bounds().Max.Y)
	})
}

func loadImage(path string) image.Image {
	imageFile, _ := os.Open("../testdata/font_1/" + path)
	defer imageFile.Close()
	img, _, _ := image.Decode(imageFile)
	return img
}
