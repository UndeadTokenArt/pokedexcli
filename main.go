package main

import (
	"time"

	"github.com/undeadtokenart/pokedexcli/internal/player"
	"github.com/undeadtokenart/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient           pokeapi.Client
	nextLocationAreaURL     *string
	previousLocationAreaURL *string
	caughtPokemon           map[string]pokeapi.Pokemon
	playerXP                int
	playerInv               player.Inventory
}

func main() {
	client := player.Inventory{
		Gold:      100,
		Pokeballs: 10,
	}

	cfg := config{
		pokeapiClient: pokeapi.NewClient(time.Hour),
		caughtPokemon: make(map[string]pokeapi.Pokemon),
		playerXP:      0,
		playerInv:     client,
	}

	startrepl(&cfg)

}
