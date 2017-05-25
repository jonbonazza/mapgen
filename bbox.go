package mapgen

import "github.com/pzsz/voronoi"

type BBox struct {
	*voronoi.BBox
	Width, Height float64
}

func NewBBox(xl, xr, yt, yb float64) *BBox {
	bb := voronoi.NewBBox(xl, xr, yt, yb)
	return &BBox{
		BBox: &bb,
		Width: xr - xl,
		Height: yb - yt,
	}
}
