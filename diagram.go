package mapgen

import (
	"github.com/pzsz/voronoi"
	"github.com/pzsz/voronoi/utils"
)

type Diagram struct {
	*voronoi.Diagram
	Center voronoi.Vertex
}

func NewDiagram(bbox voronoi.BBox, count, relaxPasses int) *Diagram {
	sites := utils.RandomSites(bbox, count)
	sites = relax(sites, bbox, relaxPasses)
	return &Diagram{
		Diagram: voronoi.ComputeDiagram(sites, bbox, true),
		Center:  center(bbox),
	}
}

func center(bbox voronoi.BBox) voronoi.Vertex {
	dx := bbox.Xr - bbox.Xl
	dy := bbox.Yb - bbox.Yt
	return voronoi.Vertex{
		X: dx / 2,
		Y: dy / 2,
	}
}

func relax(sites []voronoi.Vertex, bbox voronoi.BBox, passes int) []voronoi.Vertex {
	for i := 0; i < passes; i++ {
		diagram := voronoi.ComputeDiagram(sites, bbox, true)
		sites = utils.LloydRelaxation(diagram.Cells)
	}
	return sites
}
