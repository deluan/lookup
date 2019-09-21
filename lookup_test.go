package lookup

import (
	_ "image/png"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLookupAll(t *testing.T) {
	Convey("Given an Color image and a template to look for", t, func() {
		img := loadImageColor("testdata/cyclopst1.png")
		template := loadImageColor("testdata/cyclopst3.png")

		Convey("When searching", func() {
			pp, _ := Color(img, template, 0.9)
			Convey("It finds the template", func() {
				So(pp, ShouldHaveLength, 1)
				So(pp[0].X, ShouldEqual, 21)
				So(pp[0].Y, ShouldEqual, 7)
				So(pp[0].G, ShouldBeGreaterThan, 0.9)
			})
		})
	})
	Convey("Given an Gray Scale image and a template to look for", t, func() {
		img := loadImageGray("testdata/cyclopst1.png")
		template := loadImageGray("testdata/cyclopst3.png")

		Convey("When searching", func() {
			pp, _ := GrayScale(img, template, 0.9)
			Convey("It finds the template", func() {
				So(pp, ShouldHaveLength, 1)
				So(pp[0].X, ShouldEqual, 21)
				So(pp[0].Y, ShouldEqual, 7)
				So(pp[0].G, ShouldBeGreaterThan, 0.9)
			})
		})
	})
}
