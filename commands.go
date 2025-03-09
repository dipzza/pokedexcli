package main

import (
	"fmt"
	"os"

	"github.com/dipzza/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

type config struct {
	nextURL string
	prevURL string
}

var cliCommands map[string]cliCommand

func init() {
	cliCommands = map[string]cliCommand{
		"help": {
			name:					"help",
			description:	"Displays a help message",
			callback:			commandHelp,
		},
		"explore": {
			name:					"explore <area_name>",
			description:	"Displays Pokémons in an area",
			callback:			commandExplore,
		},
		"map": {
			name:					"map",
			description:	"Displays the next 20 locations in the map",
			callback:			commandMap,
		},
		"mapb": {
			name:					"mapb",
			description:	"Displays the previous 20 locations in the map",
			callback:			commandMapb,
		},
		"exit": {
			name:					"exit",
			description:	"Exit the Pokedex",
			callback:			commandExit,
		},
	}
}

func commandHelp(*config, ...string) error {
	fmt.Print(
		"Welcome to the Pokedex!\n",
		"Usage:\n\n",
	)

	for _, value := range cliCommands {
		fmt.Println(value.name + ":", value.description)
	}

	return nil
}

func commandExplore(c *config, params ...string) error {
	if len(params) == 0 {
		fmt.Println("Missing area name. Usage: explore <area_name>")
	}

	fmt.Println("Exploring " + params[0] + "...")
	locationArea, err := pokeapi.GetLocationArea(params[0])
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, encounter := range locationArea.PokemonEncounters {
		fmt.Println("-", encounter.Pokemon.Name)
	}

	return nil
}

func commandMap(c *config, _ ...string) error {
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

func commandMapb(c *config, _ ...string) error {
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


func commandExit(*config, ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}