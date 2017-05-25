package mapgen

import (
	"github.com/jonbonazza/mapgen/pmrng"
	"github.com/pzsz/voronoi"
	"github.com/pzsz/voronoi/utils"
)

const (
	NumRelaxIterations = 2
)

type Map struct {
	centers []*Center
}

func NewMap(bbox voronoi.BBox, numPoints int) *Map {
	verts := generateRandomVerts(bbox, numPoints)
	diagram := voronoi.ComputeDiagram(relaxVerts(verts, bbox), bbox, true)
	return &Map{

	}
}

func (m *Map) buildGraph(verts []voronoi.Vertex, diagram *voronoi.Diagram) {
	centers := make(map[*voronoi.Vertex]*Center)
	for i, vertex := range verts {
		center := &Center{
			Index:    i,
			Location: &vertex,
			Neighbors: make([]*Center, 0),
			Borders:  make([]*Edge, 0),
			Corners: make([]*Corner, 0),
		}
		m.centers = append(m.centers, center)
		centers[&vertex] = center
	}
	cornerMap := make(map[int][]*Corner)
	for _, edge := range diagram.Edges {
		e := &Edge{
			Index: len(diagram.Edges),
		}
	}
	corner := findCorner()
}

func findCorner(vert *voronoi.Vertex, cornerMap map[int][]*Corner) *Corner {
	for bucket := int(vert.X)-1; bucket <= int(vert.X)+1; bucket++ {
		for _, corner := range cornerMap[bucket] {
			dx := vert.X - corner.Location.X
			dy := vert.Y - corner.Location.Y
			if dx*dx + dy*dy < 1e-6 {
				return  corner
			}
		}
	}
	return nil
}

func relaxVerts(verts []voronoi.Vertex, bbox voronoi.BBox) []voronoi.Vertex {
	for i := 0; i < NumRelaxIterations; i++ {
		diagram := voronoi.ComputeDiagram(verts, bbox, true)
		verts = utils.LloydRelaxation(diagram.Cells)
	}
	return verts
}
