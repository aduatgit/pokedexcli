package main

import (
	"errors"
	"fmt"
)

func commandMapf(cfg *config, args ...string) error {
	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocations)
	if err != nil {
		return err
	}

	cfg.nextLocations = locationsResp.Next
	cfg.previousLocations = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapb(cfg *config, args ...string) error {
	if cfg.previousLocations == nil {
		return errors.New("you are on page 1")
	}
	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.previousLocations)
	if err != nil {
		return err
	}

	cfg.nextLocations = locationsResp.Next
	cfg.previousLocations = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}
