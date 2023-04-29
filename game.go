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
	"fmt"

	"github.com/gen2brain/raylib-go/raylib"
)


func drawGameScene(level int) (bool, int) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	log.Printf("Started new level: %d", level)

	initChans()
	defer closeChans()

	cQueue := newCustomerQueue()
	tStack := newTrayStack()

	balance := 0
	trayCount := 0
	cookedKumpirs := 0

	// Push new customers into the queue.
	go generateCustomers(ctx, levelConfig[level].spawnCustomers)
	// Handle orders, take some time to prepare each.
	go ovenCook(ctx)

	for {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		handleControlKeyPress()
		if gOpt == Pause {
			rl.DrawText("Paused.", 320, 240, 50, rl.DarkGray)
			rl.DrawText("Press Esc to continue.", 120, 300, 50, rl.DarkGray)
			rl.EndDrawing()
			continue
		}

		drawInfo(level, balance)
		drawCustomerQueue(cQueue.customers)
		drawTrayArea(tStack.trays)
		drawGameControls()

		rl.DrawTexture(TableTex, 0, 0, rl.White)
		drawCookedKumpirs(cookedKumpirs)
		if ovenStatus == 1 {
			rl.DrawText("Preparing order...", 10, 600, 30, rl.DarkGray)
		}
		select {

		// New customer arrived at the diner.
		case c := <-arrivingCustomers:
			// Animate addition to queue.
			log.Printf("Adding to queue: %d", c.Id)
			cQueue.add(c)

		// Order is ready to serve to customer.
		case <-ovenDone:
			cookedKumpirs = 3
			ct := cQueue.head()
			serveOrder(ct, &balance)

			// Customer will spend some time eating now.
			go customerEat(ctx, ct)

			// Remove customer from the queue now that we served
			// the dish.
			log.Printf("Removing from queue: %d", ct.Id)
			cQueue.remove()

		// Customer is done eating.
		case c := <-doneCustomers:
			log.Printf("Has left tray: %d", c.Id)
			cookedKumpirs--
			if tStack.length() == 5 {
				log.Printf("Tray space overflow!: %d", c.Id)
				rl.EndDrawing()
				return false, balance
			}
			tStack.add(c.Id)

		default:
			if trayCount == levelConfig[level].spawnCustomers {
				log.Println("Level done")
				rl.EndDrawing()
				return true, balance
			}

			// The player is busy preparing the order,
			// don't respond to key press
			if ovenStatus == 1 {
				rl.EndDrawing()
				continue
			}

			if tStack.length() > 0 && rl.IsKeyPressed(rl.KeyT) {
				tStack.remove()
				trayCount += 1
			}

			if cQueue.length() >= 1 {
				ovenCookOnPress()
			}
		}

		rl.EndDrawing()
	}
}

func drawCookedKumpirs(amount int) {
	if amount == 0 {
		return
	}
	texture := rl.LoadTextureFromImage(CookedImg[amount])
	rl.DrawTexture(texture, 520, 750, rl.White)
}

func drawCustomerQueue(q []Customer) {
	if len(q) != 0 {
		//for _, i := range q[0].KumpirOrder{
		/*rl.ImageDraw(PotatoImg, ToppingsImg[i].Img,
		rl.NewRectangle(0, 0, float32(PotatoImg.Width), float32(PotatoImg.Height)),
		rl.NewRectangle(0, 0, float32(PotatoImg.Width), float32(PotatoImg.Height)), rl.White)
		}*/
		tex := rl.LoadTextureFromImage(q[0].KumpirTex)
		rl.DrawTexture(tex, 200, 200, rl.White)

	}
	for i, v := range q {
		rl.DrawCircleV(rl.Vector2{400, float32(400 - 90 * i)}, 40, rl.Maroon)
		rl.DrawText(fmt.Sprintf("%v", v.KumpirOrder), 372, int32(400 - 90 * i - 15), 25, rl.LightGray)
	}
}

func drawTrayArea(s []int) {
	rl.DrawRectangle(550, 30, 100, 200, rl.Yellow)
	for i, _ := range s {
		rl.DrawRectangle(560, int32(200 - i * 40), 80, 20, rl.DarkGray)
	}

}

func drawInfo(l, b int) {
	rl.DrawText(fmt.Sprintf("Day: %d", l + 1), 10, 10, 30, rl.DarkGray)
	rl.DrawText(fmt.Sprintf("Balance: %d", b), 10, 40, 30, rl.DarkGray)
}

func drawGameControls() {
	rl.DrawRectangle(580, 470, 300, 240, rl.Green)
	rl.DrawText("K - Serve Kebab", 585, 475, 20, rl.DarkGray)
	rl.DrawText("P - Serve Pizza", 585, 500, 20, rl.DarkGray)
	rl.DrawText("H - Serve Hamburger", 585, 525, 20, rl.DarkGray)
	rl.DrawText("T - Empty tray", 585, 550, 20, rl.DarkGray)
	rl.DrawText("Esc - Pause", 585, 575, 20, rl.DarkGray)
	rl.DrawText("Q - Quit", 585, 600, 20, rl.DarkGray)

}
