package lookup

import (
	_ "image/png"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFontSymbol(t *testing.T) {
	Convey("Given an image and a template to look for", t, func() {
		img := loadImageGray("font_1/0.png")
		fs := NewFontSymbol("0", img)
		So(fs.Width, ShouldEqual, img.Bounds().Max.X)
		So(fs.Height, ShouldEqual, img.Bounds().Max.Y)
	})
}
