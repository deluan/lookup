package lookup

import (
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
			})
		})
	})
}

func loadImage(path string) image.Image {
	imageFile, _ := os.Open(path)
	defer imageFile.Close()
	img, _, _ := image.Decode(imageFile)
	return img
}
