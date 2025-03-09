package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/dipzza/pokedexcli/internal/pokeapi"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	locationsConfig := config{
		nextURL: pokeapi.LocationAreaEndpoint,
		prevURL: "",
	}

	for {
		fmt.Print("Pokedex > ")

		scanner.Scan()
		words := cleanInput(scanner.Text())

		if len(words) > 0 {
			command, ok := cliCommands[words[0]]
			if !ok {
				fmt.Println("Unknown command:", words[0])
				continue
			}

			err := command.callback(&locationsConfig, words[1:]...)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println()
			}
		}
	}
}

func cleanInput(text string) []string {
	lowerText := strings.ToLower(text)
	words := strings.Fields(lowerText)
	return words
}
