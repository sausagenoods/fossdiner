package main

import (
	"context"
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

func generateCustomers(ctx context.Context, customerSpawnAmount int) {
	id := 1
	menu := []MenuOption{Kebab, Pizza, Hamburger}
	for {
		if gOpt == Pause {
			time.Sleep(500 * time.Millisecond)
			continue
		}
		if id > customerSpawnAmount {
			return
		}
		d := time.Duration(rand.Intn(7) + 1) * time.Second
		gen := Customer{
			Order: menu[rand.Intn(3)],
			Id: id,
		}

		select {
		case <-ctx.Done():
			return

		case <-time.After(d):
			arrivingCustomers <- gen
			id++
		}
	}
}

func customerEat(ctx context.Context, c Customer) {
	// Waste some time
	d := time.Duration(rand.Intn(5) + 1) * time.Second
	select {
	case <-ctx.Done():
		return
	case <-time.After(d):
		doneCustomers <- c
	}
}
