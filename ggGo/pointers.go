package main

import "fmt"

type Player struct {
	health int
}

func (player *Player) takeDamageFromExplosion(dmg int) {
	fmt.Println("player is taking dmg: ")
	player.health -= dmg

}

func main() {
	player := &Player{
		health: 100,
	}

	fmt.Println("before explosion: %v+\n", player)
	player.takeDamageFromExplosion(50)
	fmt.Println("after explosion: %v+\n", player)
}
