package common

import (
	_ "image/png"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSArray(t *testing.T) {
	Convey("Given a SArray", t, func() {
		array, _ := NewSArray(4, 4, []float64{
			5, 2, 5, 2,
			3, 6, 3, 6,
			5, 2, 5, 2,
			3, 6, 3, 6,
		})

		Convey("It should throw an error when accessing an invalid coordinate", func() {
			err := array.Set(-1, 0, 33)
			So(err.Error(), ShouldContainSubstring, "bad coordinate")
			err = array.Set(0, -1, 33)
			So(err.Error(), ShouldContainSubstring, "bad coordinate")
		})

		Convey("When I use it as a base for a new SArray", func() {
			newArray := NewSArrayFromBase(array)
			Convey("The new SArray should have the original SArray as its base", func() {
				So(newArray.base, ShouldEqual, array)
			})
			Convey("It should have the same dimensions as base", func() {
				So(newArray.Width(), ShouldEqual, array.Width())
				So(newArray.Height(), ShouldEqual, array.Height())
			})
		})
	})

	Convey("StepThrough calls the Step method for each pixel", t, func() {
		test := &testSArray{}
		sarray, _ := NewSArray(2, 2, []float64{1, 2, 3, 4})
		test.initBase(sarray)
		test.called = []int{0, 0, 0, 0}

		stepThrough(test)

		for _, v := range test.called {
			if v != 1 {
				So(v, ShouldEqual, 1)
			}
		}
	})

}

type testSArray struct {
	basicSArray
	called []int
}

func (ta *testSArray) Step(x, y int) {
	ta.called[ta.pos(x, y)] += 1
}
