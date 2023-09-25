package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startrepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(">")

		scanner.Scan()
		text := scanner.Text()
		cleaned := cleanInput(text)

		if len(cleaned) == 0 {
			continue
		}
		command := cleaned[0]
		switch command {
		case "help":
			fmt.Println("Welcome to the pokedex help menu!")
			fmt.Println("here are your available commands")
			fmt.Println("- Help")
			fmt.Println("- Exit")
			fmt.Println("")
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("Invalid Command")
		}
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
