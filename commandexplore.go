package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}

	name := args[0]
	localEncounters, err := cfg.pokeapiClient.GetLocation(name)
	if err != nil {
		return err
	}
	fmt.Printf("\nExploring %s...\n", name)
	fmt.Println("Found Pokemon:")
	for _, v := range localEncounters.PokemonEncounters {
		fmt.Printf("- %s\n", v.Pokemon.Name)
	}
	return nil
}
