package canton

import (
	. "github.com/tobiashort/ch/coord"
)

type Canton struct {
	Name     string
	Abbr     string
	Polygons [][]Coord
}
