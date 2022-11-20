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
