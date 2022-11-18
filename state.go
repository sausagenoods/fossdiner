package main

import (
	"os"

	"github.com/gen2brain/raylib-go/raylib"
)

type GameState int

const (
	MenuScreen GameState = 0
	InGame GameState = 1
)

var gState GameState

type GameOpt int

const (
	Default GameOpt = 0
	Pause GameOpt = 1
	Preparing GameOpt = 2
)

var gOpt GameOpt

func handleControlKeyPress() {
	if rl.IsKeyPressed(rl.KeyEscape) {
		// Toggle Pause/Unpause states with XOR
		gOpt ^= 1
	}
	if rl.IsKeyPressed(rl.KeyQ) {
		quitGame()
	}
}

func quitGame() {
	// TODO: Implement a proper deinit function
	rl.CloseWindow()
	os.Exit(0)
}
