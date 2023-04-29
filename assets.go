package main

import (
	"strings"
	"os"
	"fmt"
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
var CookedImg = make(map[int]*rl.Image)
var PotatoImg = rl.LoadImage("assets/img/potato.png")
var TableTex rl.Texture2D

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

	for i := 1; i < 4; i++ {
		CookedImg[i] = rl.LoadImage(fmt.Sprintf("assets/img/cooked-%d.png", i))
	}

	tableImg := rl.LoadImage("assets/img/table.png")
	TableTex = rl.LoadTextureFromImage(tableImg)
}
