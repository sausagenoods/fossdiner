package main

import (
	"math/rand"
	"time"
	"log"
	"fmt"

	"github.com/gen2brain/raylib-go/raylib"
)

var balance int

var arrivingCustomers chan string
var eatingCustomers chan string

// For signalling tray addition to the stack
var doneCustomers chan struct{}


func generateCustomers() {
	menu := []string{"K", "P", "H"}
	for {
		time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
		log.Println("Added new customer")
		arrivingCustomers <- menu[rand.Intn(2)]
	}
}

func main() {
	arrivingCustomers = make(chan string)
	eatingCustomers = make(chan string)

	var cQueue []string

	rl.InitWindow(800, 640, "Fossdiner")
	rl.SetTargetFPS(60)

	// Push new customers into the queue
	go generateCustomers()

	balance = 0
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		drawBalance()
		// draw buttons
		drawCustomerQueue(cQueue)

		select {
		case c := <-arrivingCustomers:
			// animate addition to queue
			cQueue = append(cQueue, c)
		default:
			if (len(cQueue) < 1) {
				rl.EndDrawing()
				continue
			}
			handleServeKeys(&cQueue)
		}
		rl.EndDrawing()
	}
	rl.CloseWindow()
}

func drawCustomerQueue(cQueue []string) {
	for i, v := range cQueue {
		rl.DrawCircleV(rl.Vector2{400, float32(400 - 70 * i)}, 30, rl.Maroon)
		rl.DrawText(v, 400, int32(400 - 70 * i), 20, rl.LightGray)
	}
}

func handleServeKeys(cQueue *[]string) {
	first := (*cQueue)[0]
	correctDish := false

	if rl.IsKeyPressed(rl.KeyK) {
		if first == "K" { correctDish = true }
	} else if rl.IsKeyPressed(rl.KeyH) {
		if first == "H" { correctDish = true }
	} else if rl.IsKeyPressed(rl.KeyP) {
		if first == "P" { correctDish = true}
	} else {
		// No key press event to process
		return
	}

	// Implement custom delay while serving
	// Animate clock and play tick sound
	// time.Sleep(5 * time.Second) <- blocks,
	// use channels instead for prepared dishes.

	if correctDish {
		*cQueue = (*cQueue)[1:]
		log.Println("Removed from queue:", first)
		balance += 10
		log.Println("+Balance:", balance)
		return
	}

	balance -= 10
	log.Println("-Balance:", balance)

}

func drawBalance() {
	rl.DrawText(fmt.Sprintf("%d", balance), 190, 200, 20, rl.LightGray)
}
