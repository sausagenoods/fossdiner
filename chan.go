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

func initChans() {
	arrivingCustomers = make(chan Customer)
	eatingCustomers = make(chan Customer)
	doneCustomers = make(chan Customer)
	ovenLoad = make(chan struct{})
	ovenDone = make(chan struct{})
}

func closeChans() {
	close(arrivingCustomers)
	close(eatingCustomers)
	close(doneCustomers)
	close(ovenLoad)
	close(ovenDone)
}
