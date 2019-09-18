package common

import (
	_ "image/png"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestIntegralImage(t *testing.T) {
	Convey("Given a SArray", t, func() {
		original, _ := NewSArray(4, 4, []float64{
			5, 2, 5, 2,
			3, 6, 3, 6,
			5, 2, 5, 2,
			3, 6, 3, 6,
		})
		var sum float64 = 0
		for _, v := range original.array {
			sum += v
		}

		array := NewIntegralImageFromBase(original)

		Convey("It should calculate the right sigma value, == sum(all values)", func() {
			So(array.Sigma(), ShouldEqual, sum)
		})

		Convey("It should calculate the right mean value, == Sigma/Size", func() {
			So(array.Mean(), ShouldEqual, 4)
		})
	})
}
