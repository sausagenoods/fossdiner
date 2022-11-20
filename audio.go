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

import "github.com/gen2brain/raylib-go/raylib"

var tickSnd rl.Sound
var cashSnd rl.Sound
var wrongSnd rl.Sound

func initAudio() {
	rl.InitAudioDevice()
	tickSnd = rl.LoadSound("assets/audio/ticking.wav")
	cashSnd = rl.LoadSound("assets/audio/cashier.mp3")
	wrongSnd = rl.LoadSound("assets/audio/wrong-order.mp3")
}

func deinitAudio() {
	rl.UnloadSound(cashSnd)
	rl.UnloadSound(tickSnd)
	rl.CloseAudioDevice()
}
