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
		commandName := cleanInput(reader.Text())
		if len(commandName) == 0 {
			continue
		}

		cmd, exists := getCommands()[commandName]
		if exists {
			err := cmd.callback(cfg)
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

func cleanInput(s string) string {
	s = strings.ToLower(s)
	a := strings.Fields(s)
	return a[0]
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
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
	}
}
