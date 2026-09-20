package main

import (
	"fmt"
	_ "time"

	gui "github.com/ShuffleKamori/go-practice/GUI"
	backend "github.com/ShuffleKamori/go-practice/backend"
)

func main() {
	fmt.Println("Start!")
	go backend.Start_backend()
	if !gui.Start_gui() {
		fmt.Println("GUI failed initialize!")
	}
}
