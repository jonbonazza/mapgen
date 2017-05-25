package mapgen

import (
	"github.com/pzsz/voronoi"
)

type Edge struct {
	Index int
	D0, D1 *Center	// Delaunay edge
	V0, V1 *Corner  // Voronoi edge
	Midpoint *voronoi.Vertex // Halfway between v0, v1
	River int	// Volume of water, or 0
}

type Corner struct {
	Index int
	Location *voronoi.Vertex
	Ocean bool
	Water bool	// Lake or ocean
	Coast bool	// Touches ocean and land polygons
	Border bool	// At the edge of the map
	Elevation float64	// 0.0-1.0
	Moisture float64        // 0.0-1.0

	Touches []*Center
	Protrudes []*Edge
	Adjacent []*Corner

	River int 	// 0 if no river, or volume of water in river
	Downslopes *Corner	// adjacent corner most downhill
	Watershed *Corner	// coastal corner or null
	WatershedSize int
}

type Center struct {
	Index     int
	Location  *voronoi.Vertex
	Water     bool	// Lake or ocean
	Ocean     bool
	Coast     bool	// Land polygon touching an ocean
	Border    bool	// At the edge of the map
	Biome     string	// Biome type
	Elevation float64	// 0.0-1.0
	Moisture  float64	// 0.0-1.0
	Neighbors []*Center
	Borders   []*Edge
	Corners   []*Corner
}
