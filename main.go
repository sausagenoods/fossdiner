package main

import (
	"fmt"

	"github.com/gen2brain/raylib-go/raylib"
	"github.com/gen2brain/raylib-go/raygui"
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

	gState = MenuScreen
	// Three levels (day): 0, 1, 2
	level := 0
	for !rl.WindowShouldClose() {
		switch(gState) {
		case MenuScreen:
			rl.BeginDrawing()
			rl.ClearBackground(rl.RayWhite)
			rl.DrawText("Press the button to start.", 300, 180, 20, rl.DarkGray)
			if raygui.Button(rl.NewRectangle(400, 280, 80, 20), "START") {
				gState = InGame
			}
			rl.EndDrawing()
		case InGame:
			if fin, balance := drawGameScene(level); fin {
				for !rl.IsKeyPressed(rl.KeyEnter) {
					rl.BeginDrawing()
					rl.ClearBackground(rl.RayWhite)
					rl.DrawText("Day cleared!", 190, 160, 20, rl.DarkGray)
					rl.DrawText(fmt.Sprintf("balance: %d, day: %d", balance, level + 1), 190, 200, 20, rl.DarkGray)
					rl.DrawText("Press enter to start next day.", 190, 240, 20, rl.DarkGray)
					rl.EndDrawing()
				}
				level += 1
			} else {
				for !rl.IsKeyPressed(rl.KeyEnter) {
					rl.BeginDrawing()
					rl.ClearBackground(rl.RayWhite)
					rl.DrawText("You lose!", 190, 160, 20, rl.DarkGray)
					rl.DrawText(fmt.Sprintf("balance: %d, day: %d", balance, level + 1), 190, 200, 20, rl.DarkGray)
					rl.DrawText("Press enter to repeat same day.", 190, 240, 20, rl.DarkGray)
				}
			}
		}
	}

	deinitAudio()

	rl.CloseWindow()
}
