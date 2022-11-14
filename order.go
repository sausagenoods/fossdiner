package main

import (
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

func kitchenPrepareOrders() {
	for {
		select {
		case o := <- kitchenOrder:
			rl.PlaySound(tickSnd)
			orderStatus = 1
			// Take up to 5 seconds to prepare the order.
			time.Sleep(time.Duration(rand.Intn(7) + 1) * time.Second)
			orderStatus = 0
			// Send a signal now that the order is ready.
			orderReady <- o

			rl.PauseSound(tickSnd)
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
