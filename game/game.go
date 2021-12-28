package game

import (
	raylib "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	Width  int32
	Height int32
	Title  string

	TargetFps       int32
	BackgroundColor raylib.Color
}

func CreateGame(width, height int32, title string, bgColor raylib.Color) *Game {
	return &Game{
		Width:           width,
		Height:          height,
		Title:           title,
		TargetFps:       60,
		BackgroundColor: bgColor,
	}
}

func (g *Game) Run() {
	raylib.InitWindow(g.Width, g.Height, g.Title)
	raylib.SetTargetFPS(g.TargetFps)

	for !raylib.WindowShouldClose() {
		raylib.BeginDrawing()

		raylib.ClearBackground(g.BackgroundColor)

		raylib.EndDrawing()
	}

	raylib.CloseWindow()
}
