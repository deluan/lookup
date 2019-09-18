package lookup

import (
	"image"
	_ "image/png"
	"os"
	"testing"

	"github.com/deluan/lookup/common"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLookupAll(t *testing.T) {
	Convey("Given an image and a template to look for", t, func() {
		img := loadImage("cyclopst1.png")
		template := loadImage("cyclopst3.png")

		Convey("When searching in RGB", func() {
			pp, _ := LookupAllGrey(img, template, 0.9)
			Convey("It finds the template", func() {
				So(pp, ShouldHaveLength, 1)
				So(pp[0].X, ShouldEqual, 21)
				So(pp[0].Y, ShouldEqual, 7)
				So(pp[0].G, ShouldBeGreaterThan, 0.9)
			})
		})
	})
}

func TestMultiplyAndSum(t *testing.T) {
	Convey("Given a two arrays", t, func() {
		a1, _ := common.NewSArray(2, 2, []float64{1, 2, 3, 4})
		a2, _ := common.NewSArray(2, 2, []float64{1, 2, 3, 4})

		Convey("It sums all resulting pixels", func() {
			sum := multiplyAndSum(a1, 0, 0, a2)
			So(sum, ShouldEqual, 1+4+9+16)
		})
	})
}

var (
	benchImg         = loadImage("cyclopst1.png")
	benchTemplate    = loadImage("cyclopst3.png")
	benchImgBin      = common.NewImageBinaryGrey(benchImg)
	benchTemplateBin = common.NewImageBinaryGrey(benchTemplate)
)

func BenchmarkLookupAll(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = lookupAll(benchImgBin, benchTemplateBin, 0.9)
	}
}

func BenchmarkMultiplyAndSum(b *testing.B) {
	b.StopTimer()
	imgBin := common.NewImageBinaryGrey(benchImg)
	templateBin := common.NewImageBinaryGrey(benchTemplate)
	ci := imgBin.Channels()[0]
	ct := templateBin.Channels()[0]
	b.StartTimer()

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		multiplyAndSum(ci.ZeroMeanImage(), 0, 0, ct.ZeroMeanImage())
	}
}

func loadImage(path string) image.Image {
	imageFile, _ := os.Open("testdata/" + path)
	defer imageFile.Close()
	img, _, _ := image.Decode(imageFile)
	return img
}
