package main

import (
	"math/rand"
	"time"
	"log"
	"fmt"

	"github.com/gen2brain/raylib-go/raylib"
)

type MenuOption rune

const (
	Kebab MenuOption = 'K'
	Pizza MenuOption = 'P'
	Hamburger MenuOption = 'H'
)

type Customer struct {
	Order MenuOption
	Id int
}

var tickSnd rl.Sound
var cashSnd rl.Sound
var wrongSnd rl.Sound

var balance int

var arrivingCustomers chan Customer
var eatingCustomers chan Customer

// For signalling tray addition to the stack
var doneCustomers chan Customer

// Signal the kitchen that an order needs to be prepared.
var kitchenOrder chan MenuOption
// For signalling when order is ready to serve.
var orderReady chan MenuOption

func generateCustomers() {
	menu := []MenuOption{Kebab, Pizza, Hamburger}
	for {
		gen := Customer{
			Order: menu[rand.Intn(2)],
			Id: rand.Intn(999),
		}

		time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
		arrivingCustomers <- gen
	}
}

func kitchenPrepareOrders() {
	for {
		select {
		case o := <- kitchenOrder:
			rl.PlaySound(tickSnd)

			// Take up to 5 seconds to prepare the order.
			time.Sleep(time.Duration(rand.Intn(4) + 1) * time.Second)
			// Send a signal now that the order is ready.
			orderReady <- o

			rl.PauseSound(tickSnd)
		}
	}
}

func main() {
	// Initialize assets, channels...
	arrivingCustomers = make(chan Customer)
	eatingCustomers = make(chan Customer)
	doneCustomers = make(chan Customer)
	kitchenOrder = make(chan MenuOption)
	orderReady = make(chan MenuOption)

	var cQueue []Customer
	var tStack []int

	rl.InitWindow(800, 640, "Fossdiner")
	rl.SetTargetFPS(60)

	rl.InitAudioDevice()
	tickSnd = rl.LoadSound("assets/audio/ticking.wav")
	cashSnd = rl.LoadSound("assets/audio/cashier.mp3")
	wrongSnd = rl.LoadSound("assets/audio/wrong-order.mp3")

	// Push new customers into the queue.
	go generateCustomers()
	// Handle orders, take some time to prepare each.
	go kitchenPrepareOrders()

	balance = 0
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		drawBalance()
		// draw buttons
		drawCustomerQueue(cQueue)

		select {
		// New customer arrived at the diner.
		case c := <-arrivingCustomers:
			// animate addition to queue
			log.Printf("Adding to queue: %c-%d", c.Order, c.Id)
			cQueue = append(cQueue, c)

		// Order is ready to serve to customer.
		case c := <-orderReady:
			serveOrder(cQueue[0], c)
			// Remove customer from the queue now that we served
			// the dish.
			log.Printf("Removing from queue: %c-%d", cQueue[0].Order, cQueue[0].Id)
			cQueue = cQueue[1:]

		// Customer is done eating.
		case c := <-doneCustomers:
			tStack = append(tStack, c.Id)

		default:
			if (len(cQueue) < 1) {
				rl.EndDrawing()
				continue
			}
			placeKitchenOrderOnKeyPress()
		}

		rl.EndDrawing()
	}

	rl.UnloadSound(cashSnd)
	rl.UnloadSound(tickSnd)

	rl.CloseAudioDevice()

	rl.CloseWindow()
}

func drawCustomerQueue(cQueue []Customer) {
	for i, v := range cQueue {
		rl.DrawCircleV(rl.Vector2{400, float32(400 - 70 * i)}, 30, rl.Maroon)
		rl.DrawText(fmt.Sprintf("%c-%d", v.Order, v.Id), 400, int32(400 - 70 * i), 15, rl.LightGray)
	}
}

func placeKitchenOrderOnKeyPress() {
	if rl.IsKeyPressed(rl.KeyK) {
		kitchenOrder <- Kebab
	} else if rl.IsKeyPressed(rl.KeyH) {
		kitchenOrder <- Hamburger
	} else if rl.IsKeyPressed(rl.KeyP) {
		kitchenOrder <- Pizza
	}
}

func serveOrder(c Customer, m MenuOption) {
	// Does what we're serving match the customer's order?
	if c.Order == m {
		rl.PlaySound(cashSnd)
		balance += 10
		log.Println("+Balance:", balance)
		return
	}
	rl.PlaySound(wrongSnd)
	balance -= 10
	log.Println("-Balance:", balance)
}

func drawBalance() {
	rl.DrawText(fmt.Sprintf("%d", balance), 190, 200, 20, rl.LightGray)
}
