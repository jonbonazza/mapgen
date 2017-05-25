package mapgen

import (
	"math/rand"
	"github.com/pzsz/voronoi/utils"
)

func (m *Map) generateTopography() {
	u := m.Unit
	for i, cell := range m.Diagram.Cells {
		d := utils.Distance(cell.Site, m.Diagram.Center)
		n := m.noise.Noise2D(
			cell.Site.X/(m.BoundingBox.Width/4),
			cell.Site.Y/(m.BoundingBox.Height/4),
		)
		e := getElevation(m, d, n)

		c := &Cell{
			Index: i,
			CenterDistance: d,
			NoiseLevel: n,
			Elevation: e,
			Land: e >= 0,
			Site: cell.Site,
		}
		if c.Land {
			// Make sure edges of the map are water
			if (cell.Site.X < u*0.5 || cell.Site.X > float64(m.BoundingBox.Width)-u*0.5) ||
				(cell.Site.Y < u/1.5 || cell.Site.Y > float64(m.BoundingBox.Height)-u/1.5) ||
				(cell.Site.Y < u/3 || cell.Site.Y > float64(m.BoundingBox.Height)-u/3) {
				c.Land = false
				c.Elevation = -1.5 * c.NoiseLevel
			}
		}
		if c.Land {
			if d < u*3.3 {
				c.Elevation += 0.3
			}
			if d < u*2.3 {
				c.Elevation += 0.6
			}
			if d < u*1.3 {
				c.Elevation += 0.9
			}
			// Add some lakes
			if c.NoiseLevel < -0.3 {
				c.Elevation = c.NoiseLevel
			}
		}
		colorCell(c)
		m.Cells = append(m.Cells, c)
	}
}

func colorCell(c *Cell) {
	if c.Land {
		switch {
		case c.Elevation > 7:
			c.FillColor = Green7
			c.StrokeColor = Green7
		case c.Elevation > 6.1:
			c.FillColor = Green6
			c.StrokeColor = Green6
		case c.Elevation > 4.8:
			c.FillColor = Green5
			c.StrokeColor = Green5
		case c.Elevation > 3.1:
			c.FillColor = Green4
			c.StrokeColor = Green4
		case c.Elevation > 2.4:
			c.FillColor = Green3
			c.StrokeColor = Green3
		case c.Elevation > 1.5:
			c.FillColor = Green2
			c.StrokeColor = Green2
		case c.Elevation < -0.6:
			c.FillColor = Blue1
			c.StrokeColor = Blue1
		case c.Elevation < -0.4:
			c.FillColor = Blue0
			c.StrokeColor = Blue0
		case c.Elevation < 0:
			c.FillColor = Yellow1
			c.StrokeColor = Yellow1
		default:
			c.FillColor = Green1
			//c.StrokeColor = Green2
			c.StrokeColor = Green1
		}
	} else {
		switch {
		case c.Elevation < -6:
			c.FillColor = Blue7
			c.StrokeColor = Blue7
		case c.Elevation < -5:
			c.FillColor = Blue6
			c.StrokeColor = Blue6
		case c.Elevation < -4:
			c.FillColor = Blue5
			c.StrokeColor = Blue5
		case c.Elevation < -3:
			c.FillColor = Blue4
			c.StrokeColor = Blue4
		case c.Elevation < -2:
			c.FillColor = Blue3
			c.StrokeColor = Blue3
		case c.Elevation < -1:
			c.FillColor = Blue2
			c.StrokeColor = Blue2
		default:
			c.FillColor = Blue1
			//c.StrokeColor = Blue2
			c.StrokeColor = Blue1
		}
	}
}

func (m *Map) getNeighbors(cell *Cell) []*Cell {
	neighbors := make([]*Cell, 0)
	for i, dcell := range m.Diagram.Cells {
		for _, hedge := range m.Diagram.Cells[cell.Index].Halfedges {
			if utils.EdgeIndex(dcell, hedge.Edge) >= 0 {
				neighbors = append(neighbors, m.Cells[i])
			}
		}
	}
	return neighbors
}

func getElevation(m *Map, distance, noise float64) (e float64) {
	e = 1.8 + noise
	e -= distance / m.Unit / 3.75
	if e > 0 {
		e += 1 + float64(rand.Int63n(2))
		if e > 1.5 && rand.Intn(3) < 2 {
			e += 0.5 + rand.Float64()
		}
		if e > 3 {
			e += 1.5 + rand.Float64()
		}
	}
	return
}