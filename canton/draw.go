package canton

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	. "github.com/tobiashort/ch/coord"
)

func Draw(canton Canton) {
	for _, polygon := range canton.Polygons {
		for i := range polygon {
			startCoord := polygon[i]
			var endCoord Coord
			if i+1 == len(polygon) {
				endCoord = polygon[0]
			} else {
				endCoord = polygon[i+1]
			}
			startVec := CoordToXY(startCoord)
			endVec := CoordToXY(endCoord)
			rl.DrawLineEx(startVec, endVec, 1, rl.White)
		}
	}
}
