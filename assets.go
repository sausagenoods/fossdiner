package main

import (
	"strings"
	"os"
	"log"
	"strconv"
	"path/filepath"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Topping struct {
	Img *rl.Image
	Priority int
}

var ToppingsImg = make(map[string]Topping)
var PotatoImg = rl.LoadImage("assets/img/potato.png")

func initImgAssets() {

	toppingFiles, err := os.ReadDir("assets/img/toppings")
	if err != nil {
		log.Fatal(err)
	}
	for _, t := range toppingFiles {
		fn := t.Name()
		log.Println(fn)
		tName := strings.Split(strings.Split(fn, ".")[0], "-")
		log.Println(tName[0], tName[1])
		priority, err := strconv.Atoi(tName[1])
		if err != nil {
			log.Fatal(err)
		}
		topping := Topping{
			Img: rl.LoadImage(filepath.Join("assets/img/toppings", fn)),
			Priority: priority,
		}
		ToppingsImg[tName[0]] = topping
	}
}
