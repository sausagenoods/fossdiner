package main

import (
	"fmt"

	"github.com/gen2brain/raylib-go/raylib"
)

func main() {
	// Initialize assets, channels...
	arrivingCustomers = make(chan Customer)
	eatingCustomers = make(chan Customer)
	doneCustomers = make(chan Customer)
	kitchenOrder = make(chan MenuOption)
	orderReady = make(chan MenuOption)

	rl.InitWindow(800, 640, "Fossdiner")
	rl.SetTargetFPS(60)

	initAudio()

	gState = InGame
	// Three levels (day): 0, 1, 2
	level := 0
	for !rl.WindowShouldClose() {
		switch(gState) {
		case InGame:
			if fin, balance := drawGameScene(level); fin {
				for !rl.IsKeyPressed(rl.KeyEnter) {
					rl.BeginDrawing()
					rl.ClearBackground(rl.RayWhite)
					rl.DrawText("Day cleared!", 190, 160, 20, rl.LightGray)
					rl.DrawText(fmt.Sprintf("balance: %d, level: %d", balance, level), 190, 200, 20, rl.LightGray)
					rl.DrawText("Press enter to start next day.", 190, 240, 20, rl.LightGray)
					rl.EndDrawing()
				}
				level++
			}
		}
	}

	deinitAudio()

	rl.CloseWindow()
}
