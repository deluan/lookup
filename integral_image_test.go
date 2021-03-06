package lookup

import (
	_ "image/png"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestIntegralImage(t *testing.T) {
	Convey("Given a Integral Image", t, func() {
		pixels := []uint8{
			5, 2, 5, 2,
			3, 6, 3, 6,
			5, 2, 5, 2,
			3, 6, 3, 6,
		}
		grayImage := newGrayImage(4, 4, pixels)
		sum := 0
		for _, v := range pixels {
			sum += int(v)
		}

		Convey("When I calculate the Integral Image of it", func() {
			integral := newIntegralImage(grayImage)

			Convey("Then its sigma is the sum of all pixels", func() {
				So(integral.sigma(integral.pix, 0, 0, 3, 3), ShouldEqual, sum)
			})
			Convey("And its mean is sigma/size", func() {
				So(integral.mean, ShouldEqual, sum/len(pixels))
			})
		})

	})
}

func BenchmarkNewIntegralImage(b *testing.B) {
	b.StopTimer()
	grayImage := newGrayImage(4, 4, []uint8{
		5, 2, 5, 2,
		3, 6, 3, 6,
		5, 2, 5, 2,
		3, 6, 3, 6,
	})
	b.StartTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		newIntegralImage(grayImage)
	}
}
