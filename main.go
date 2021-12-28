package main

import (
	"game-template/game"

	raylib "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth  = 800
	screenHeight = 600
	title        = "Template"
)

func init() {
	raylib.SetConfigFlags(raylib.FlagVsyncHint)
	raylib.SetConfigFlags(raylib.FlagMsaa4xHint)
	raylib.SetConfigFlags(raylib.FlagWindowResizable)
}

func main() {
	g := game.CreateGame(screenWidth, screenHeight, title, raylib.White)

	g.Run()
}
