package lookup

import (
	"image"
	_ "image/png"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNumerator(t *testing.T) {
	Convey("Given two imageBinary structs", t, func() {
		a1 := image.NewGray(image.Rect(0, 0, 2, 2))
		a2 := image.NewGray(image.Rect(0, 0, 2, 2))
		b1 := newImageBinaryChannel(a1, gray)
		b2 := newImageBinaryChannel(a2, gray)
		b1.zeroMeanImage = []float64{1, 2, 3, 4}
		b2.zeroMeanImage = []float64{1, 2, 3, 4}

		Convey("It sums all zeroMean images pixels", func() {
			sum := numerator(b1, b2, 0, 0)
			So(sum, ShouldEqual, 1+4+9+16)
		})
	})
}

var (
	benchImg         = loadImageColor("testdata/cyclopst1.png")
	benchTemplate    = loadImageColor("testdata/cyclopst3.png")
	benchImgBin      = newImageBinary(benchImg)
	benchTemplateBin = newImageBinary(benchTemplate)
)

func BenchmarkLookupAll(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = lookupAll(benchImgBin, benchTemplateBin, 0.9)
	}
}

func BenchmarkNumerator(b *testing.B) {
	ci := benchImgBin.channels[0]
	ct := benchTemplateBin.channels[0]

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		numerator(ci, ct, 0, 0)
	}
}
