package canton

import rl "github.com/gen2brain/raylib-go/raylib"

type Coordinates struct {
	North float32
	West  float32
}

type Canton struct {
	Name     string
	Abbr     string
	Polygons [][]Coordinates
}

func coordinatesToXY(coordinates Coordinates) rl.Vector2 {
	const zoom float32 = 70

	width := float32(rl.GetScreenWidth()) * zoom
	height := float32(rl.GetScreenHeight()) * zoom

	offsetCoordinates := Coordinates{North: 5.7, West: 48.0}
	offsetX := (offsetCoordinates.North + 180) * (width / 360)
	offsetY := (90 - offsetCoordinates.West) * (height / 180)

	x := (coordinates.North+180)*(width/360) - offsetX
	y := (90-coordinates.West)*(height/180) - offsetY
	return rl.Vector2{X: x, Y: y}
}

func DrawCanton(canton Canton) {
	for _, polygon := range canton.Polygons {
		for i := range polygon {
			startCoord := polygon[i]
			var endCoord Coordinates
			if i+1 == len(polygon) {
				endCoord = polygon[0]
			} else {
				endCoord = polygon[i+1]
			}
			startVec := coordinatesToXY(startCoord)
			endVec := coordinatesToXY(endCoord)
			rl.DrawLineEx(startVec, endVec, 1, rl.White)
		}
	}
}

func DrawAll() {
	DrawCanton(AG)
	DrawCanton(AI)
	DrawCanton(AR)
	DrawCanton(BE)
	DrawCanton(BL)
	DrawCanton(BS)
	DrawCanton(FR)
	DrawCanton(GE)
	DrawCanton(GL)
	DrawCanton(GR)
	DrawCanton(JU)
	DrawCanton(LU)
	DrawCanton(NE)
	DrawCanton(NW)
	DrawCanton(OW)
	DrawCanton(SG)
	DrawCanton(SH)
	DrawCanton(SO)
	DrawCanton(SZ)
	DrawCanton(TG)
	DrawCanton(TI)
	DrawCanton(UR)
	DrawCanton(VD)
	DrawCanton(VS)
	DrawCanton(ZG)
	DrawCanton(ZH)
}
