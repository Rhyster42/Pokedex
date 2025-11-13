package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Rhyster42/Pokedex/internal/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
	caughtPokemon    map[string]pokeapi.PokemonInfo
}

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		var commandArgs string
		if len(words) > 1 {
			commandArgs = strings.Join(words[1:], " ")
		}

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg, commandArgs)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Printf("Unknown command\n")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Shows the next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Shows the previous 20 locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore <location_name>",
			description: "Shows Pokemon of current location. Requires location argument.",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon_name>",
			description: "Attempt to catch Pokemon. If succesful, adds Pokemon to Pokedex. Requires pokemon argument.",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon_name>",
			description: "Inspect pokemon in Pokedex. Requires Pokemon argument",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "List Pokemon in Pokedex",
			callback:    commandPokedex,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
