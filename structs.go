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
