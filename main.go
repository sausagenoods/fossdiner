package main

import (
	"math/rand"
	"time"
	"log"

	"github.com/gen2brain/raylib-go/raylib"
)

var cChan chan string

func generateCustomers() {
	menu := []string{"K", "P", "H"}
	for {
		time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
		log.Println("Added new customer")
		cChan <- menu[rand.Intn(2)]
	}
}

func main() {
//	var customerQueue []int
	cChan = make(chan string)
	cQueue := make([]string, 0)

	rl.InitWindow(800, 640, "raylib [core] example - basic window")
	rl.SetTargetFPS(60)

	// Push new customers into the queue
	go generateCustomers()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		// draw buttons
		drawCustomerQueue(cQueue)

		select {
		case c := <-cChan:
			// animate addition to queue
			cQueue = append(cQueue, c)
			break
		default:
			break
		}

		if len(cQueue) == 0 {
			continue
		}

		rl.EndDrawing()

		first := cQueue[0]
		if rl.IsKeyPressed(rl.KeyK) {
			if first == "K" {
				cQueue = cQueue[1:]
				continue
			}
			rl.DrawText("Wrong!", 190, 200, 20, rl.LightGray)
			// Animate K button down
		} else if rl.IsKeyPressed(rl.KeyH) {
			if first == "H" {
				cQueue = cQueue[1:]
				continue
			}
			rl.DrawText("Wrong!", 190, 200, 20, rl.LightGray)

		} else if rl.IsKeyPressed(rl.KeyP) {
			if first == "P" {
				cQueue = cQueue[1:]
				continue
			}
			rl.DrawText("Wrong!", 190, 200, 20, rl.LightGray)
		}
	}
	rl.CloseWindow()
}

func drawCustomerQueue(cQueue []string) {
	for i, v := range cQueue {
		rl.DrawCircleV(rl.Vector2{400, float32(400 - 70 * i)}, 30, rl.Maroon)
		rl.DrawText(v, 400, int32(400 - 70 * i), 20, rl.LightGray)
	}
}
