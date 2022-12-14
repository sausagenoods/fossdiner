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
	"log"
	"math/rand"
	"time"

	"github.com/gen2brain/raylib-go/raylib"
)

type MenuOption rune

const (
	Kebab MenuOption = 'K'
	Pizza MenuOption = 'P'
	Hamburger MenuOption = 'H'
)

// Signal the kitchen that an order needs to be prepared.
var kitchenOrder chan MenuOption
// For signalling when order is ready to serve.
var orderReady chan MenuOption

// 0 = none/done, 1 = preparing
var orderStatus int
func placeKitchenOrderOnKeyPress() {
	if rl.IsKeyPressed(rl.KeyK) {
		kitchenOrder <- Kebab
	} else if rl.IsKeyPressed(rl.KeyH) {
		kitchenOrder <- Hamburger
	} else if rl.IsKeyPressed(rl.KeyP) {
		kitchenOrder <- Pizza
	}
}

func kitchenPrepareOrders(ctx context.Context) {
	defer log.Println("kitchenPrepareOrders exit")
	for {
		if gOpt == Pause {
			time.Sleep(500 * time.Millisecond)
			continue
		}
		select {
		case o := <- kitchenOrder:
			d := time.Duration(rand.Intn(7) + 1) * time.Second
			rl.PlaySound(tickSnd)
			orderStatus = 1

			select {
			case <-time.After(d):
				rl.PauseSound(tickSnd)
				orderStatus = 0
				orderReady <- o
				//case <-ctx.Done():
			case <-ctx.Done():
				rl.PauseSound(tickSnd)
				orderStatus = 0
				return
			}

		case <-ctx.Done():
			rl.PauseSound(tickSnd)
			orderStatus = 0
			return
		}
	}
}

func serveOrder(c Customer, m MenuOption, b *int) {
	// Does what we're serving match the customer's order?
	if c.Order == m {
		rl.PlaySound(cashSnd)
		*b += 10
		log.Println("+Balance:", *b)
		return
	}
	rl.PlaySound(wrongSnd)
	*b -= 10
	log.Println("-Balance:", *b)
}
