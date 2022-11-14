package main

import "github.com/gen2brain/raylib-go/raylib"

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
	for !rl.WindowShouldClose() {
		switch(gState) {
		case InGame:
			drawGameScene()
		}
	}

	deinitAudio()

	rl.CloseWindow()
}
