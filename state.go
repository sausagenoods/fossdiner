package main

type GameState int

const (
	MenuScreen GameState = 0
	InGame GameState = 1
)

var gState GameState

type GameOpt int

const (
	Default GameOpt = 0
	Pause GameOpt = 1
	Debug GameOpt = 2
)

var gOpt GameOpt
