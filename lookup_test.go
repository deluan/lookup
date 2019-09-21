package lookup

import (
	_ "image/png"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLookup(t *testing.T) {
	Convey("Given a GrayScale image", t, func() {
		img := loadImageGray("testdata/cyclopst1.png")
		l := NewLookupGrayScale(img)

		Convey("When I search using a grayscale template", func() {
			template := loadImageGray("testdata/cyclopst3.png")

			Convey("It finds the template", func() {
				pp, _ := l.FindAll(template, 0.9)
				So(pp, ShouldHaveLength, 1)
				So(pp[0].X, ShouldEqual, 21)
				So(pp[0].Y, ShouldEqual, 7)
				So(pp[0].G, ShouldBeGreaterThan, 0.9)
			})
		})

		Convey("When I search using a color template", func() {
			template := loadImageColor("testdata/cyclopst3.png")

			Convey("It finds the template", func() {
				pp, _ := l.FindAll(template, 0.9)
				So(pp, ShouldHaveLength, 1)
				So(pp[0].X, ShouldEqual, 21)
				So(pp[0].Y, ShouldEqual, 7)
				So(pp[0].G, ShouldBeGreaterThan, 0.9)
			})
		})
	})

	Convey("Given a Color image", t, func() {
		img := loadImageColor("testdata/cyclopst1.png")
		l := NewLookup(img)

		Convey("When I search using a color template", func() {
			template := loadImageColor("testdata/cyclopst3.png")

			Convey("It finds the template", func() {
				pp, _ := l.FindAll(template, 0.9)
				So(pp, ShouldHaveLength, 1)
				So(pp[0].X, ShouldEqual, 21)
				So(pp[0].Y, ShouldEqual, 7)
				So(pp[0].G, ShouldBeGreaterThan, 0.9)
			})
		})
		Convey("When I search using a grayscale template", func() {
			template := loadImageGray("testdata/cyclopst3.png")

			Convey("It returns an error", func() {
				_, err := l.FindAll(template, 0.9)
				println(err.Error())
				So(err, ShouldNotBeNil)
			})
		})

	})
}
