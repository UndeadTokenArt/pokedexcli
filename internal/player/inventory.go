package player

import (
	"errors"
	"fmt"
)

type Inventory struct {
	gold      int
	pokeballs int
}

func (playerInv *Inventory) ThrowPokeBall(num int) error {
	if playerInv.pokeballs < 1 {
		return errors.New("player has no more pokeballs to throw")
	}
	playerInv.pokeballs -= num
	fmt.Printf("Player has %v Pokeballs in thier backpack left\n", playerInv.pokeballs)
	return nil
}

func (playerInv *Inventory) purchasPokeBalls(num int) error {
	cost := num * 30
	if playerInv.gold <= cost {
		return errors.New("you dont have enough gold")
	}
	playerInv.pokeballs += num
	playerInv.gold -= cost
	fmt.Printf("Player now has purchased %v pokeballs\n Purchase cost: %v gold \nThey have %v Pokeballs in their backpack\n", num, cost, playerInv.pokeballs)
	return nil
}

func (playerInv *Inventory) AddGoldToPlayer(num int) {
	playerInv.gold += num
}
