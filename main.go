package main

import (
	"fmt"
	"time"
)

type player struct {
	name   string
	health int
	posX   float32
	posY   float32
	posZ   float32
}

func main() {
	player1 := player{name: "Shuffle", health: 17, posX: 100.1, posY: 100.1, posZ: 100.1}
	player2 := new(player)
	fmt.Println(player1)
	fmt.Println(player2)
	fmt.Println(player1.name)
	fmt.Println(player1.health)
	fmt.Println(player1.posX)
	fmt.Println(player1.posY)
	time.Sleep(5 * time.Second)
}
