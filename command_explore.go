package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, locationArea string) error {
	if locationArea == "" {
		return errors.New("no location argument, please provide location to explore")
	}

	pokemonResp, err := cfg.pokeapiClient.LocationDeepInfo(locationArea)
	if err != nil {
		return err
	}

	for _, pokemon := range pokemonResp.PokemonEncounters {
		fmt.Println(pokemon.Pokemon.Name)
	}
	return nil
}
