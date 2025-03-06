package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

type cliCommand struct {
	name 				string
	description string
	callback		func() error
}

var cliCommands map[string]cliCommand

func init() {
	cliCommands = map[string]cliCommand{
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for true {
		fmt.Print("Pokedex > ")

		scanner.Scan()
		words := cleanInput(scanner.Text())

		if len(words) > 0 {
			command, ok := cliCommands[words[0]]
			if !ok {
				fmt.Println("Unknown command:", words[0])
				continue
			}

			err := command.callback()
			if err != nil {
				fmt.Println("Error:", err)
			}
		}
	}
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println(`
Welcome to the Pokedex!
Usage:
`)

	for key, value := range cliCommands {
		fmt.Println(key, ":", value.description)
	}

	return nil
}

func cleanInput(text string) []string {
	lowerText := strings.ToLower(text)
	words := strings.Fields(lowerText)
	return words
}

