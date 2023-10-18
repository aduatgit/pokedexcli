package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/aduatgit/pokedexcli/internal/pokeapi"
)

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
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		cmd, exists := getCommands()[commandName]
		if exists {
			err := cmd.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Printf("%s is not a valid command\n", commandName)
			continue
		}

	}
}

func cleanInput(s string) []string {
	s = strings.ToLower(s)
	a := strings.Fields(s)
	return a
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

type config struct {
	caughtPokemon     map[string]pokeapi.Pokemon
	pokeapiClient     pokeapi.Client
	nextLocations     *string
	previousLocations *string
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"mapf": {
			name:        "mapf",
			description: "Displays the names of the next 20 location areas in the Pokemon world",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the names of the previous 20 location areas in the Pokemon world",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore <location>",
			description: "Displays possible pokemon encounters of a location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon>",
			description: "Attempts to catch a pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon>",
			description: "Attempts to inspect a pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Returns all the pokemon in your pokedex",
			callback:    commandPokedex,
		},
	}
}
