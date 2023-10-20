package player

import (
	"errors"
	"fmt"
)

type Inventory struct {
	Gold      int
	Pokeballs int
}

func (playerInv *Inventory) ThrowPokeBall(num int) error {
	if playerInv.Pokeballs < 1 {
		return errors.New("player has no more pokeballs to throw")
	}
	playerInv.Pokeballs -= num
	fmt.Printf("Player has %v Pokeballs in thier backpack left\n", playerInv.Pokeballs)
	return nil
}

func (playerInv *Inventory) PurchasPokeBalls(num int) error {
	cost := num * 200
	if playerInv.Gold <= cost {
		return errors.New("you dont have enough gold")
	}
	playerInv.Pokeballs += num
	playerInv.Gold -= cost
	fmt.Printf("Player now has purchased %v pokeballs\n Purchase cost: %v gold \nThey have %v Pokeballs in their backpack\n", num, cost, playerInv.Pokeballs)
	return nil
}

func (playerInv *Inventory) AddGoldToPlayer(num int) {
	playerInv.Gold += num
}
