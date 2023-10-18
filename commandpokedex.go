package main

import "fmt"

func commandPokedex(cfg *config, args ...string) error {
	fmt.Println("Your Pokedex:")
	for _, v := range cfg.caughtPokemon {
		fmt.Printf(" - %s\n", v.Name)
	}
	return nil
}
