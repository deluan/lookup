package lookup

import (
	"github.com/deluan/lookup/common"
	"image"
	_ "image/png"
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLookupAll(t *testing.T) {
	Convey("Given an image and a template to look for", t, func() {
		img := loadImage("examples/cyclopst1.png")
		template := loadImage("examples/cyclopst3.png")

		Convey("When searching in RGB", func() {
			pp, _ := LookupAll(img, template, 0.9)
			Convey("It finds the template", func() {
				So(pp, ShouldHaveLength, 1)
				So(pp[0].X, ShouldEqual, 21)
				So(pp[0].Y, ShouldEqual, 7)
				So(pp[0].G, ShouldBeGreaterThan, 0.9)
			})
		})
	})
}

var (
	benchImg      = loadImage("examples/cyclopst1.png")
	benchTemplate = loadImage("examples/cyclopst3.png")
)

func BenchmarkLookupAll(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = LookupAll(benchImg, benchTemplate, 0.9)
	}
}

func BenchmarkNumerator(b *testing.B) {
	b.StopTimer()
	imgBin := common.NewImageBinaryGrey(benchImg)
	templateBin := common.NewImageBinaryGrey(benchTemplate)
	ci := imgBin.Channels()[0]
	ct := templateBin.Channels()[0]
	b.StartTimer()

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		common.NewImageMultiply(ci.ZeroMeanImage(), 0, 0, ct.ZeroMeanImage())
	}
}

func loadImage(path string) image.Image {
	imageFile, _ := os.Open(path)
	defer imageFile.Close()
	img, _, _ := image.Decode(imageFile)
	return img
}
