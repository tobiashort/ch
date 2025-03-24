package coord_test

import (
	"math"
	"testing"

	. "github.com/tobiashort/cantons/coord"
)

func TestCoordToXYtoCoord(t *testing.T) {
	c := Coord{North: 7.9564886, West: 47.455536}
	xy := CoordToXY(c)
	c1 := XYtoCoord(xy)
	nd := math.Abs(float64(c.North - c1.North))
	wd := math.Abs(float64(c.West - c1.West))
	if nd > 0.001 || wd > 0.001 {
		t.Errorf("N:%f->%f,W:%f->%f", c.North, c1.North, c.West, c1.West)
	}
}
