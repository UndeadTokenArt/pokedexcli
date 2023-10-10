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
	playerInv.pokeballs += num
	cost := num * 30
	if playerInv.gold <= cost {
		return errors.New("You dont have enough Gold")
	}
	playerInv.gold -= cost
	fmt.Printf("Player now has purchased %v pokeballs\n Purchase cost: %v gold \nThey have %v Pokeballs in their backpack\n", num, cost, playerInv.pokeballs)
}

func (playerInv *Inventory) AddGoldToPlayer(num int) error {
	if playerInv.gold < 1 {
		return errors.New("player has no gold")
	}
	playerInv.gold += num
	return nil
}
