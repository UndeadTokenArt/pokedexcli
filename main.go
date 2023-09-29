package main

import (
	"time"

	"github.com/undeadtokenart/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient           pokeapi.Client
	nextLocationAreaURL     *string
	previousLocationAreaURL *string
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(time.Hour),
	}

	startrepl(&cfg)

}
