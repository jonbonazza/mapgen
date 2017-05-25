package mapgen

import (
	"github.com/pzsz/voronoi"
	"image/color"
	"image"
	"math"
	"math/rand"
	"github.com/llgcode/draw2d/draw2dimg"
)

type Cell struct {
	Index int
	CenterDistance float64
	NoiseLevel float64
	Elevation float64
	Land bool
	Site voronoi.Vertex
	FillColor color.RGBA
	StrokeColor color.RGBA
	Neighbors []*Cell
}

type Map struct {
	BoundingBox *BBox
	Unit    float64
	Cells   []*Cell
	Diagram *Diagram
	noise   *Noise
}

func NewMap(bbox *BBox, siteCount, relaxPasses int, unit float64) *Map {
	m := &Map{
		BoundingBox: bbox,
		Unit:    unit,
		Cells:   make([]*Cell, 0),
		Diagram: NewDiagram(*bbox.BBox, siteCount, relaxPasses),
		noise:   NewNoise(rand.Int63n(int64(bbox.Width*bbox.Height))),
	}
	m.generateTopography()
	return m
}

func (m *Map) Image() image.Image {
	img := image.NewRGBA(image.Rect(0, 0, int(m.BoundingBox.Width), int(m.BoundingBox.Height)))
	ctx := draw2dimg.NewGraphicContext(img)
	defer ctx.Close()
	ctx.SetLineWidth(1.2)
	for i, cell := range m.Diagram.Cells {
		ctx.SetFillColor(m.Cells[i].FillColor)
		ctx.SetStrokeColor(m.Cells[i].StrokeColor)
		for _, hedge := range cell.Halfedges {
			a := hedge.GetStartpoint()
			b := hedge.GetEndpoint()
			ctx.MoveTo(a.X, a.Y)
			ctx.LineTo(b.X, b.Y)
		}
		ctx.FillStroke()
	}
	return img
}
