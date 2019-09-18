package common

import (
	_ "image/png"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestImageMultiply(t *testing.T) {
	Convey("Given a ImageMultiply array", t, func() {
		a1, _ := NewSArray(2, 2, []float64{1, 2, 3, 4})
		a2, _ := NewSArray(2, 2, []float64{1, 2, 3, 4})

		im := NewImageMultiply(a1, 0, 0, a2)

		Convey("It multiply all pixels", func() {
			So(im.array, ShouldResemble, []float64{1, 4, 9, 16})
		})

		Convey("It sums all resulting pixels", func() {
			So(im.Sum(), ShouldEqual, 1+4+9+16)
		})
	})
}
