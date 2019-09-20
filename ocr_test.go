package lookup

import (
	_ "image/png"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestOCR(t *testing.T) {
	Convey("Given an OCR object", t, func() {
		ocr := NewOCR(0.7)

		Convey("When I try to load an invalid font directory", func() {
			err := ocr.loadFont("testdata/NON_EXISTENT")

			Convey("It returns an error", func() {
				So(err.Error(), ShouldContainSubstring, "no such file or directory")
			})
		})

		Convey("When I load a valid font on it", func() {
			err := ocr.loadFont("testdata/font_1")

			Convey("It loads the fonts successfully", func() {
				So(err, ShouldBeNil)
			})

			Convey("It stores the font family", func() {
				So(ocr.fontFamilies, ShouldContainKey, "font_1")
				So(ocr.fontFamilies, ShouldHaveLength, 1)
				So(ocr.fontFamilies["font_1"], ShouldHaveLength, 13)
			})
		})
	})
}
