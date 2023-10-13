package main

import (
	"errors"
	"fmt"
)

func callShop(cfg *config, args ...string) error {

	resp, err := cfg.pokeapiClient.ListItems(cfg.nextItem)
	if err != nil {
		return err
	}

	fmt.Println("Location areas:")
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}
	cfg.nextItem = resp.Next
	cfg.previousItem = resp.Previous
	return nil
}

func callShopb(cfg *config, args ...string) error {
	if cfg.previousItem == nil {
		return errors.New("you are on the first page")
	}

	resp, err := cfg.pokeapiClient.ListItems(cfg.previousItem)
	if err != nil {
		return err
	}

	fmt.Println("Location areas:")
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}
	cfg.nextItem = resp.Next
	cfg.previousItem = resp.Previous
	return nil
}
