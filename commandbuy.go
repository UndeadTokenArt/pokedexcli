package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func callbackBuy(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no Item chosen")
	}
	itemName := args[0]

	itemDetails, err := cfg.pokeapiClient.GetItemDetails(itemName)
	if err != nil {
		return err
	}
	fmt.Printf("Item costs:  %v:\n", itemDetails.Cost)
	fmt.Printf("You have %v gold\n", cfg.playerInv.Gold)
	fmt.Printf("Would you like to buy %s?\n", itemDetails.Name)
	fmt.Print(">")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(">")

		scanner.Scan()
		text := scanner.Text()
		cleaned := cleanInput(text)

		if len(cleaned) == 0 {
			continue
		}

		if cleaned[0] == "yes" {
			pokeballs := cfg.playerInv.Pokeballs
			gold := cfg.playerInv.Gold
			cost := itemDetails.Cost
			if gold < cost {
				return errors.New("not enough gold to make purchase")
			}
			gold -= cost
			pokeballs += 1
			fmt.Printf("You now have %v pokeballs and %v Gold\n", pokeballs, gold)
			return nil
		} else {
			return errors.New("declined to purchase")
		}
	}
}
