package main

func initChans() {
	arrivingCustomers = make(chan Customer)
	eatingCustomers = make(chan Customer)
	doneCustomers = make(chan Customer)
	kitchenOrder = make(chan MenuOption)
	orderReady = make(chan MenuOption)
}

func closeChans() {
	close(arrivingCustomers)
	close(eatingCustomers)
	close(doneCustomers)
	close(kitchenOrder)
	close(orderReady)
}
