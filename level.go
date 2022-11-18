package main

type levelParams struct {
	spawnCustomers int
}

var levelConfig = [3]levelParams {
	{spawnCustomers: 7,}, {spawnCustomers: 10,}, {spawnCustomers: 15,},
}
