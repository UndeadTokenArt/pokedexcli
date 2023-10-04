package main

import (
	"fmt"
)

func callbackPokedex(cfg *config, args ...string) error {
	if len(cfg.caughtPokemon) < 1 {
		fmt.Println("You have not caught any pokemon")
	} else {
		fmt.Println("Pokemon in Pokedex")
		for _, pokemon := range cfg.caughtPokemon {
			fmt.Printf("- %s\n", pokemon.Name)
		}
	}
	return nil
}
