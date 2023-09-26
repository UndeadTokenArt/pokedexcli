package main

import (
	"fmt"
	"log"

	"github.com/undeadtokenart/pokedexcli/internal/pokeapi"
)

func main() {
	// startrepl()

	pokeapiClient := pokeapi.NewClient()

	resp, err := pokeapiClient.ListLocationAreas()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp)

}
