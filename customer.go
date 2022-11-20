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
