package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startrepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(">")

		scanner.Scan()
		text := scanner.Text()
		cleaned := cleanInput(text)

		if len(cleaned) == 0 {
			continue
		}
		commandName := cleaned[0]
		args := []string{}
		if len(cleaned) > 1 {
			args = cleaned[1:]
		}

		availableCommands := getCommands()

		command, ok := availableCommands[commandName]
		if !ok {
			fmt.Println("Invalid Command")
			continue
		}
		err := command.callback(cfg, args...)
		if err != nil {
			fmt.Println(err)
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Prints the help menu",
			callback:    callBackHelp,
		},
		"exit": {
			name:        "exit",
			description: "Ends the program",
			callback:    callbackExit,
		},
		"map": {
			name:        "map",
			description: "Lists Next set of location Areas",
			callback:    callbackMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Lists Previous set of location Areas",
			callback:    callbackMapb,
		},
		"explore": {
			name:        "explore {location_area}",
			description: "Lists the pokemon in a location area",
			callback:    callbackExplore,
		},
		"catch": {
			name:        "catch {pokemon_name}",
			description: "Attempt to catch a pokemon and add it to your pokedex",
			callback:    callbackCatch,
		},
		"inspect": {
			name:        "inspect {pokemon_name}",
			description: "See some details about selected Pokemon",
			callback:    callbackInspect,
		},
		"pokedex": {
			name:        "Pokedex",
			description: "Shows a list of all the pokemon caught",
			callback:    callbackPokedex,
		},
		"shop": {
			name:        "Shop",
			description: "Get a list of 20 Items to shop for",
			callback:    callShop,
		},
		"shopb": {
			name:        "Shop",
			description: "Get a list of 20 Items to shop for",
			callback:    callbackShopb,
		},
		"buy": {
			name:        "buy",
			description: "Buy an Item from the shop",
			callback:    callbackBuy,
		},
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
