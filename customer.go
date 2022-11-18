package main

import (
	"time"
	"math/rand"
)

type Customer struct {
	Order MenuOption
	Id int
}

var arrivingCustomers chan Customer
var eatingCustomers chan Customer

// For signalling tray addition to the stack
var doneCustomers chan Customer

func generateCustomers(customerSpawnAmount int) {
	id := 1
	menu := []MenuOption{Kebab, Pizza, Hamburger}
	for {
		if gOpt == Pause {
			time.Sleep(500 * time.Second)
			continue
		}
		if id > customerSpawnAmount {
			return
		}
		gen := Customer{
			Order: menu[rand.Intn(3)],
			Id: id,
		}

		time.Sleep(time.Duration(rand.Intn(7) + 1) * time.Second)
		arrivingCustomers <- gen
		id++
	}
}

func customerEat(c Customer) {
	// Waste some time
	time.Sleep(time.Duration(rand.Intn(5) + 1) * time.Second)
	doneCustomers <- c
}
