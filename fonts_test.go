package lookup

import (
	_ "image/png"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFontSymbol(t *testing.T) {
	Convey("When I create a fontSymbol from a image file", t, func() {
		img := loadImageGray("testdata/font_1/0.png")
		fs := NewFontSymbol("0", img)
		Convey("It loads the image as a imageBinary", func() {
			So(fs.image, ShouldHaveSameTypeAs, &imageBinary{})
			So(fs.image.width, ShouldEqual, img.Bounds().Max.X)
			So(fs.image.height, ShouldEqual, img.Bounds().Max.Y)
			So(fs.symbol, ShouldEqual, "0")
			So(fs.width, ShouldEqual, img.Bounds().Max.X)
			So(fs.height, ShouldEqual, img.Bounds().Max.Y)
		})
	})
}

func TestLoadFont(t *testing.T) {
	Convey("Given a font directory", t, func() {
		Convey("When loading the symbols", func() {
			fonts, _ := loadFont("testdata/font_1")

			Convey("It loads all font files", func() {
				So(len(fonts), ShouldEqual, 13)
			})

			Convey("It loads all symbol names correctly", func() {
				var expectedNames = []string{"/", "€", "€", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
				var actualNames []string
				for _, f := range fonts {
					actualNames = append(actualNames, f.symbol)
				}

				So(actualNames, ShouldResemble, expectedNames)
			})
		})
	})
}
