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
	"log"
	"context"
	"time"
	"math/rand"
	"sort"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var ToppingOptions = []string{"olives", "pickles", "salad", "corn", "ketchup", "mustard"}

type Customer struct {
	KumpirOrder []string
	KumpirTex *rl.Image
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

		numOfToppings := rand.Intn(5) + 1
		var tAvailable = make([]string, len(ToppingOptions))
		copy(tAvailable, ToppingOptions)
		var kumpirOrder []string
		for i := 0; i < numOfToppings; i++ {
			log.Println(len(tAvailable))
			tIndex := rand.Intn(len(tAvailable))
			kumpirOrder = append(kumpirOrder, tAvailable[tIndex])
			tAvailable = append(tAvailable[:tIndex], tAvailable[tIndex+1:]...)
		}

		// Make sure all the toppings are visible
		sort.Slice(kumpirOrder, func(i, j int) bool {
			return ToppingsImg[kumpirOrder[i]].Priority < ToppingsImg[kumpirOrder[j]].Priority
		})

		// Draw toppings over the potato
		base := rl.ImageCopy(PotatoImg)
		for _, t := range kumpirOrder {
			rl.ImageDraw(base, ToppingsImg[t].Img,
				rl.NewRectangle(0, 0, float32(PotatoImg.Width), float32(PotatoImg.Height)),
				rl.NewRectangle(0, 0, float32(PotatoImg.Width), float32(PotatoImg.Height)), rl.White)
		}

		//rl.UnloadImage(&PotatoImg)
		gen := Customer{
			Order: menu[rand.Intn(3)],
			Id: id,
			KumpirOrder: kumpirOrder,
			KumpirTex: base,
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
