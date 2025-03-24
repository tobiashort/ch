package canton

import (
	. "github.com/tobiashort/cantons/coord"
)

type Canton struct {
	Name     string
	Abbr     string
	Polygons [][]Coord
}
