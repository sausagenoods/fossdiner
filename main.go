/*
 * FOSSdiner - the FOSS diner game.
 * Copyright (C) 2022 İrem Kuyucu <siren@kernal.eu>
 *
 *  This program is free software: you can redistribute it and/or modify
 *  it under the terms of the GNU General Public License as published by
 *  the Free Software Foundation, either version 3 of the License, or
 *  (at your option) any later version.
 *
 *  This program is distributed in the hope that it will be useful,
 *  but WITHOUT ANY WARRANTY; without even the implied warranty of
 *  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *  GNU General Public License for more details.
 *
 *  You should have received a copy of the GNU General Public License
 *  along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package main

import (
	"fmt"

	"github.com/gen2brain/raylib-go/raylib"
	"github.com/gen2brain/raylib-go/raygui"
)

func main() {
	// Initialize assets, channels...
	rl.InitWindow(800, 640, "Fossdiner")
	rl.SetTargetFPS(60)
	rl.SetTraceLog(rl.LogTrace)
	initAudio()
	initImgAssets()
	raygui.SetStyleProperty(raygui.GlobalTextFontsize, 30)

	gState = MenuScreen
	// Three levels (day): 0, 1, 2
	level := 0
	for !rl.WindowShouldClose() {
		switch(gState) {
		case MenuScreen:
			rl.BeginDrawing()
			rl.ClearBackground(rl.RayWhite)
			rl.DrawText("Press the button to start.", 220, 200, 30, rl.DarkGray)
			if raygui.Button(rl.NewRectangle(350, 280, 80, 20), "START") {
				gState = InGame
			}
			rl.EndDrawing()
		case InGame:
			if fin, balance := drawGameScene(level); fin {
				for !rl.IsKeyPressed(rl.KeyEnter) && !rl.WindowShouldClose() {
					rl.BeginDrawing()
					rl.ClearBackground(rl.RayWhite)
					rl.DrawText("Day cleared!", 190, 160, 20, rl.DarkGray)
					rl.DrawText(fmt.Sprintf("balance: %d, day: %d", balance, level + 1), 190, 200, 20, rl.DarkGray)
					rl.DrawText("Press enter to start next day.", 190, 240, 20, rl.DarkGray)
					rl.EndDrawing()
				}
				level += 1
			} else if !fin {
				for !rl.IsKeyPressed(rl.KeyEnter) && !rl.WindowShouldClose() {
					rl.BeginDrawing()
					rl.ClearBackground(rl.RayWhite)
					rl.DrawText("You lose!", 190, 160, 20, rl.DarkGray)
					rl.DrawText(fmt.Sprintf("balance: %d, day: %d", balance, level + 1), 190, 200, 20, rl.DarkGray)
					rl.DrawText("Press enter to repeat same day.", 190, 240, 20, rl.DarkGray)
					rl.EndDrawing()
				}
			}
		}
	}

	deinitAudio()

	rl.CloseWindow()
}
