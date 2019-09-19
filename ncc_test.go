package lookup

import (
	"image"
	_ "image/png"
	"os"
	"testing"

	"github.com/deluan/lookup/common"
	"github.com/deluan/lookup/utils"
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
		a1 := image.NewGray(image.Rect(0, 0, 2, 2))
		a2 := image.NewGray(image.Rect(0, 0, 2, 2))
		b1 := common.NewImageBinaryChannel(a1, common.Gray)
		b2 := common.NewImageBinaryChannel(a2, common.Gray)
		b1.ZeroMeanImage = []float64{1, 2, 3, 4}
		b2.ZeroMeanImage = []float64{1, 2, 3, 4}

		Convey("It sums all zeroMean pixels", func() {
			sum := numerator(b1, b2, 0, 0)
			So(sum, ShouldEqual, 1+4+9+16)
		})
	})
}

var (
	benchImg         = loadImage("cyclopst1.png")
	benchTemplate    = loadImage("cyclopst3.png")
	benchImgBin      = common.NewImageBinary(benchImg)
	benchTemplateBin = common.NewImageBinary(benchTemplate)
)

func BenchmarkLookupAll(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = lookupAll(benchImgBin, benchTemplateBin, 0.9)
	}
}

func BenchmarkMultiplyAndSum(b *testing.B) {
	b.StopTimer()
	imgBin := common.NewImageBinary(benchImg)
	templateBin := common.NewImageBinary(benchTemplate)
	ci := imgBin.Channels[0]
	ct := templateBin.Channels[0]
	b.StartTimer()

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		numerator(ci, ct, 0, 0)
	}
}

func loadImage(path string) image.Image {
	imageFile, _ := os.Open("testdata/" + path)
	defer imageFile.Close()
	img, _, _ := image.Decode(imageFile)
	return utils.ConvertToAverageGrayScale(img)
}

func newGrayImage(width, height int, pixels []uint8) image.Image {
	grayImage := image.NewGray(image.Rect(0, 0, width, height))
	for i, v := range pixels {
		grayImage.Pix[i] = v
	}
	return grayImage
}
