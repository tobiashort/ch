package coord

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/tobiashort/ch/globals"
)

type Coord struct {
	North float32
	West  float32
}

func CoordToXY(coord Coord) rl.Vector2 {
	width := float32(globals.InitWidth) * globals.Zoom
	height := float32(globals.InitHeight) * globals.Zoom
	x := (coord.North+180)*(width/360) - globals.OffsetX
	y := (90-coord.West)*(height/180) - globals.OffsetY
	return rl.Vector2{X: x, Y: y}
}

func XYtoCoord(vec rl.Vector2) Coord {
	width := float32(globals.InitWidth) * globals.Zoom
	height := float32(globals.InitHeight) * globals.Zoom
	north := ((360 * (vec.X + globals.OffsetX)) / width) - 180
	west := 90 - ((180 * (vec.Y + globals.OffsetY)) / height)
	return Coord{North: north, West: west}
}
