package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/tobiashort/cantons/canton"
)

//go:generate go build -o canton/gen/cantongen canton/gen/cantongen.go
//go:generate canton/gen/cantongen

func main() {
	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.InitWindow(800, 600, "Cantons of Switzerland")
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		canton.DrawAll()
		rl.EndDrawing()
	}
	rl.CloseWindow()
}
