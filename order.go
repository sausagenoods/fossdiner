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
	"context"
	"math/rand"
	"time"
	"log"

	"github.com/gen2brain/raylib-go/raylib"
)

var ovenLoad chan struct{}
var ovenDone chan struct{}

var ovenStatus int
func ovenCook(ctx context.Context) {
	for {
		if gOpt == Pause {
			time.Sleep(500 * time.Millisecond)
			continue
		}
		select {
		case <- ovenLoad:
			d := time.Duration(rand.Intn(7) + 1) * time.Second
			rl.PlaySound(tickSnd)
			ovenStatus = 1

			select {
			case <-time.After(d):
				rl.PauseSound(tickSnd)
				ovenStatus = 0
				ovenDone <- struct{}{}

			case <-ctx.Done():
				rl.PauseSound(tickSnd)
				ovenStatus = 0
				return
			}

		case <-ctx.Done():
			rl.PauseSound(tickSnd)
			ovenStatus = 0
			return
		}
	}
}

func ovenCookOnPress() {
	if rl.IsKeyPressed(rl.KeyK) {
		ovenLoad <- struct{}{}
	}
}

func serveOrder(c Customer, b *int) {
	// Does what we're serving match the customer's order?
	//if c.Order == m {
		rl.PlaySound(cashSnd)
		*b += 10
		log.Println("+Balance:", *b)
	/*	return
	}
	rl.PlaySound(wrongSnd)
	*b -= 10
	log.Println("-Balance:", *b)*/
}
