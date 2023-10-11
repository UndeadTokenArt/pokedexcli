package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func callbackCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no pokemon name provided")
	}
	pokemonName := args[0]

	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	// Add the mechanic for calculating if a pokeball is there to be thrown

	threshold := 30 + cfg.playerXP
	randNum := rand.Intn(pokemon.BaseExperience)
	fmt.Println(pokemon.BaseExperience, randNum, threshold)
	if randNum > threshold {
		return fmt.Errorf("failed to catch %s", pokemonName)
	}

	cfg.caughtPokemon[pokemonName] = pokemon
	fmt.Printf("%s was Caught!\n", pokemonName)

	// add experince to player from successful catch
	XP := pokemon.BaseExperience / 3
	cfg.playerXP += cfg.playerXP + XP
	fmt.Printf("Player gained %v XP", XP)
	return nil
}
