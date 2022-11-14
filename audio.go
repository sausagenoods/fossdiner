package main

import "github.com/gen2brain/raylib-go/raylib"

var tickSnd rl.Sound
var cashSnd rl.Sound
var wrongSnd rl.Sound

func initAudio() {
	rl.InitAudioDevice()
	tickSnd = rl.LoadSound("assets/audio/ticking.wav")
	cashSnd = rl.LoadSound("assets/audio/cashier.mp3")
	wrongSnd = rl.LoadSound("assets/audio/wrong-order.mp3")
}

func deinitAudio() {
	rl.UnloadSound(cashSnd)
	rl.UnloadSound(tickSnd)
	rl.CloseAudioDevice()
}
