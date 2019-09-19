package lookup

import (
	"fmt"
	"image"
	_ "image/png"
	"os"
	"testing"

	"github.com/deluan/lookup/common"
	"github.com/deluan/lookup/utils"
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
	benchImg         = loadImageColor("cyclopst1.png")
	benchTemplate    = loadImageColor("cyclopst3.png")
	benchImgBin      = common.NewImageBinary(benchImg)
	benchTemplateBin = common.NewImageBinary(benchTemplate)
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

func loadImageColor(path string) image.Image {
	imageFile, _ := os.Open("testdata/" + path)
	defer imageFile.Close()
	img, _, _ := image.Decode(imageFile)
	return img
}

func loadImageGray(path string) image.Image {
	img := loadImageColor(path)
	return utils.ConvertToAverageGrayScale(img)
}

func newGrayImage(width, height int, pixels []uint8) image.Image {
	grayImage := image.NewGray(image.Rect(0, 0, width, height))
	for i, v := range pixels {
		grayImage.Pix[i] = v
	}
	return grayImage
}
