package main

import (
	. "github.com/gen2brain/raylib-go/raylib"
	"github.com/tobiashort/ch/canton"
	. "github.com/tobiashort/ch/coord"
	"github.com/tobiashort/ch/globals"
)

//go:generate find canton -maxdepth 1 -regex "canton/[a-z][a-z].go" -delete
//go:generate rm -f canton/draw_all.go
//go:generate go run canton/gen/gen.go

func main() {
	SetConfigFlags(FlagWindowResizable)
	InitWindow(globals.InitWidth, globals.InitHeight, globals.Title)
	for !WindowShouldClose() {
		// Zooming
		wheelDelta := GetMouseWheelMove()
		if wheelDelta != 0 {
			xy := GetMousePosition()
			coord := XYtoCoord(xy)
			globals.Zoom += wheelDelta * globals.ScrollSpeed
			if globals.Zoom < globals.MinZoom {
				globals.Zoom = globals.MinZoom
			}
			xyAfter := CoordToXY(coord)
			dx := xy.X - xyAfter.X
			dy := xy.Y - xyAfter.Y
			globals.OffsetX -= dx
			globals.OffsetY -= dy
		}

		// Paning
		if IsMouseButtonDown(MouseButtonLeft) {
			mouseDelta := GetMouseDelta()
			globals.OffsetX -= mouseDelta.X
			globals.OffsetY -= mouseDelta.Y
		}

		BeginDrawing()
		ClearBackground(Black)
		canton.DrawAll()
		EndDrawing()
	}
	CloseWindow()
}
