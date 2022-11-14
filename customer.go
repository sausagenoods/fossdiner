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
	id := 0
	menu := []MenuOption{Kebab, Pizza, Hamburger}
	for {
		if id == customerSpawnAmount {
			return
		}
		gen := Customer{
			Order: menu[rand.Intn(2)],
			Id: id,
		}

		time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
		arrivingCustomers <- gen
		id++
	}
}
