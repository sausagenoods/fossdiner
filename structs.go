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

type trayStack struct {
	trays []int
}

func newTrayStack() *trayStack {
	var t trayStack
	return &t
}

func (t *trayStack) add(val int) {
	// Add to end
	t.trays = append(t.trays, val)
}

func (t *trayStack) remove() {
	// Remove from end
	t.trays = t.trays[:len(t.trays) - 1]
}

func (t *trayStack) length() int {
	return len(t.trays)
}

type customerQueue struct {
	customers []Customer
}

func newCustomerQueue() *customerQueue {
	var c customerQueue
	return &c
}

func (c *customerQueue) add(val Customer) {
	// Add to end
	c.customers = append(c.customers, val)
}

func (c *customerQueue) remove() {
	// Delete from beginning
	c.customers = c.customers[1:]
}

func (c *customerQueue) length() int {
	return len(c.customers)
}

func (c *customerQueue) head() Customer {
	return c.customers[0]
}
