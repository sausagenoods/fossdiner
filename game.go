package main

import (
	"log"
	"fmt"

	"github.com/gen2brain/raylib-go/raylib"
)

func drawGameScene(level int) (bool, int) {
	log.Printf("Started new level: %d", level)
	var cQueue []Customer
	var tStack []int

	balance := 0
	customerCount := 0
	trayCount := 0

	// Push new customers into the queue.
	go generateCustomers(levelConfig[level].spawnCustomers)
	// Handle orders, take some time to prepare each.
	go kitchenPrepareOrders()

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
		drawCustomerQueue(cQueue)
		drawTrayArea(tStack)
		drawGameControls()

		if orderStatus == 1 {
			rl.DrawText("Preparing order...", 10, 600, 30, rl.DarkGray)
		}
		select {

		// New customer arrived at the diner.
		case c := <-arrivingCustomers:
			// Animate addition to queue.
			log.Printf("Adding to queue: %c-%d", c.Order, c.Id)
			cQueue = append(cQueue, c)

		// Order is ready to serve to customer.
		case c := <-orderReady:
			serveOrder(cQueue[0], c, &balance)

			// Customer will spend some time eating now.
			go customerEat(cQueue[0])

			// Remove customer from the queue now that we served
			// the dish.
			log.Printf("Removing from queue: %c-%d", cQueue[0].Order, cQueue[0].Id)
			cQueue = cQueue[1:]
			customerCount += 1

		// Customer is done eating.
		case c := <-doneCustomers:
			log.Printf("Has left tray: %c-%d", c.Order, c.Id)
			if len(tStack) == 5 {
				log.Printf("Tray space overflow!: %c-%d", c.Order, c.Id)
				rl.EndDrawing()
				return false, balance
			}
			tStack = append(tStack, c.Id)

		default:
			if customerCount == levelConfig[level].spawnCustomers &&
			  trayCount == levelConfig[level].spawnCustomers {
				log.Println("Level done")
				rl.EndDrawing()
				return true, balance
			}

			// The player is busy preparing the order,
			// don't respond to key press
			if orderStatus == 1 {
				rl.EndDrawing()
				continue
			}

			if len(tStack) > 0 && rl.IsKeyPressed(rl.KeyT) {
				tStack = tStack[:len(tStack) - 1]
				trayCount += 1
			}

			if (len(cQueue) >= 1) {
				placeKitchenOrderOnKeyPress()
			}
		}

		rl.EndDrawing()
	}
}

func drawCustomerQueue(q []Customer) {
	for i, v := range q {
		rl.DrawCircleV(rl.Vector2{400, float32(400 - 90 * i)}, 40, rl.Maroon)
		rl.DrawText(fmt.Sprintf("%c-%d", v.Order, v.Id), 372, int32(400 - 90 * i - 15), 25, rl.LightGray)
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
