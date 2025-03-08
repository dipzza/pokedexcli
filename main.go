package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/dipzza/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	nextURL string
	prevURL string
}

var cliCommands map[string]cliCommand

func init() {
	cliCommands = map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the next 20 locations in the map",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 locations in the map",
			callback:    commandMapb,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	locationsConfig := config{
		nextURL: pokeapi.LocationEndpoint,
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

			err := command.callback(&locationsConfig)
			if err != nil {
				fmt.Println("Error:", err)
			}
		}
	}
}

func commandExit(*config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(*config) error {
	fmt.Print(
		"Welcome to the Pokedex!\n",
		"Usage:\n\n",
	)

	for key, value := range cliCommands {
		fmt.Println(key + ":", value.description)
	}

	return nil
}

func commandMap(c *config) error {
	if c.nextURL == "" {
		fmt.Println("You already are at the end of the map.")
		return nil
	}

	locations, err := pokeapi.GetResourcePage(c.nextURL)
	if err != nil {
		return err
	}

	c.nextURL = locations.Next
	c.prevURL = locations.Previous

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapb(c *config) error {
	if c.prevURL == "" {
		fmt.Println("You already are at the beggining of the map.")
		return nil
	}

	locations, err := pokeapi.GetResourcePage(c.prevURL)
	if err != nil {
		return err
	}

	c.nextURL = locations.Next
	c.prevURL = locations.Previous

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func cleanInput(text string) []string {
	lowerText := strings.ToLower(text)
	words := strings.Fields(lowerText)
	return words
}
