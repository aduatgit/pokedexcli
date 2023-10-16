package main

import (
	"errors"
	"fmt"
)

func commandMapf(cfg *config) error {
	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocations)
	if err != nil {
		return err
	}

	cfg.nextLocations = locationsResp.Next
	cfg.previousLocations = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc)
	}
	return nil
}

func commandMapb(cfg *config) error {
	if cfg.previousLocations == nil {
		return errors.New("You are on page 1!")
	}
	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.previousLocations)
	if err != nil {
		return err
	}

	cfg.nextLocations = locationsResp.Next
	cfg.previousLocations = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc)
	}

	return nil
}
