package testdata

import (
	"testing"

	"github.com/anthonynsimon/bild/imgio"
	"github.com/anthonynsimon/bild/transform"
	"github.com/disintegration/imaging"
)

var srcImg, err = imgio.Open("desktop.png")

func BenchmarkBildResize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		transform.Resize(srcImg, 840, 525, transform.Gaussian)
	}
}

func BenchmarkImagingResize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		imaging.Resize(srcImg, 840, 525, imaging.Gaussian)
	}
}
