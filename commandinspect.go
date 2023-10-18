package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("please provide a pokemon to inspect")
	}
	pkmnName := args[0]

	pokemon, ok := cfg.caughtPokemon[pkmnName]
	if !ok {
		return errors.New("you have not caught that pokemon")
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Printf("Stats:\n")
	for _, v := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", v.Stat.Name, v.BaseStat)
	}
	fmt.Printf("Types:\n")
	for _, v := range pokemon.Types {
		fmt.Printf("  - %s\n", v.Type.Name)
	}
	return nil
}
