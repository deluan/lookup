package lookup

import (
	"fmt"
	"image"
	_ "image/png"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLookupAll(t *testing.T) {
	Convey("Given an Color image and a template to look for", t, func() {
		img := loadImageColor("cyclopst1.png")
		template := loadImageColor("cyclopst3.png")

		Convey("When searching", func() {
			pp, _ := LookupAll(img, template, 0.9)
			Convey("It finds the template", func() {
				So(pp, ShouldHaveLength, 1)
				So(pp[0].X, ShouldEqual, 21)
				So(pp[0].Y, ShouldEqual, 7)
				So(pp[0].G, ShouldBeGreaterThan, 0.9)
				fmt.Printf("Color: %v", pp[0])
			})
		})
	})
	Convey("Given an Gray Scale image and a template to look for", t, func() {
		img := loadImageGray("cyclopst1.png")
		template := loadImageGray("cyclopst3.png")

		Convey("When searching", func() {
			pp, _ := LookupAll(img, template, 0.9)
			Convey("It finds the template", func() {
				So(pp, ShouldHaveLength, 1)
				So(pp[0].X, ShouldEqual, 21)
				So(pp[0].Y, ShouldEqual, 7)
				So(pp[0].G, ShouldBeGreaterThan, 0.9)
				fmt.Printf("Gray: %v", pp[0])
			})
		})
	})

}

func TestNumerator(t *testing.T) {
	Convey("Given a two arrays", t, func() {
		a1 := image.NewGray(image.Rect(0, 0, 2, 2))
		a2 := image.NewGray(image.Rect(0, 0, 2, 2))
		b1 := newImageBinaryChannel(a1, Gray)
		b2 := newImageBinaryChannel(a2, Gray)
		b1.ZeroMeanImage = []float64{1, 2, 3, 4}
		b2.ZeroMeanImage = []float64{1, 2, 3, 4}

		Convey("It sums all zeroMean images pixels", func() {
			sum := numerator(b1, b2, 0, 0)
			So(sum, ShouldEqual, 1+4+9+16)
		})
	})
}

var (
	benchImg         = loadImageColor("cyclopst1.png")
	benchTemplate    = loadImageColor("cyclopst3.png")
	benchImgBin      = NewImageBinary(benchImg)
	benchTemplateBin = NewImageBinary(benchTemplate)
)

func BenchmarkLookupAllColor(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = lookupAll(benchImgBin, benchTemplateBin, 0.9)
	}
}

func BenchmarkMultiplyAndSum(b *testing.B) {
	ci := benchImgBin.Channels[0]
	ct := benchTemplateBin.Channels[0]

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		numerator(ci, ct, 0, 0)
	}
}
