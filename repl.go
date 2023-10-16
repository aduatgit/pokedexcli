package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex >")
		reader.Scan()
		commandName := cleanInput(reader.Text())
		if len(commandName) == 0 {
			continue
		}

		cmd, exists := getCommands()[commandName]
		if exists {
			err := cmd.callback()
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
	callback    func() error
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
	}
}
