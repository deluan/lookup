package common

import "fmt"

// Float64 array wrapper class. Holds a []float64 array and provides basic methods.
type SArray interface {
	Get(x, y int) float64
	Set(x, y int, value float64) error
	Array() []float64
	Width() int
	Height() int
	Size() int
	Step(x, y int)
}

type basicSArray struct {
	SArray
	base   SArray
	cx, cy int
	array  []float64
}

func NewSArray(cx, cy int, values ...[]float64) (*basicSArray, error) {
	if cx <= 0 || cy <= 0 {
		return nil, fmt.Errorf("bad dimensions: (%d, %d)", cx, cy)
	}

	a := &basicSArray{
		cx: cx,
		cy: cy,
	}

	if len(values) > 0 {
		a.array = values[0]
	} else {
		a.array = make([]float64, cx*cy)
	}

	return a, nil
}

func NewSArrayFromBase(base SArray) *basicSArray {
	a := &basicSArray{}
	a.initBase(base)
	return a
}

func (s *basicSArray) initBase(base SArray) {
	s.cx = base.Width()
	s.cy = base.Height()
	s.array = make([]float64, s.cx*s.cy)
	s.base = base
}

func (s *basicSArray) Width() int {
	return s.cx
}

func (s *basicSArray) Height() int {
	return s.cy
}

func (s *basicSArray) Size() int {
	return s.cx * s.cy
}

func (s *basicSArray) Get(x, y int) float64 {
	if x < 0 || y < 0 {
		return 0
	}

	return s.array[y*s.cx+x]
}

func (s *basicSArray) Set(x, y int, value float64) error {
	if x < 0 || y < 0 {
		return fmt.Errorf("bad coordinate: (%d, %d)", x, y)
	}

	s.array[y*s.cx+x] = value
	return nil
}

func (s *basicSArray) Array() []float64 {
	return s.array
}

func stepThrough(a SArray) {
	cx := a.Width()
	cy := a.Height()
	for x := 0; x < cx; x++ {
		for y := 0; y < cy; y++ {
			a.Step(x, y)
		}
	}
}
