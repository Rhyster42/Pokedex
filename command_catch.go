package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, pokemon string) error {
	if pokemon == "" {
		return errors.New("no pokemon argument, please add pokemon you are attempting to catch")
	}

	pokemonResp, err := cfg.pokeapiClient.GetPokemon(pokemon)
	if err != nil {
		return err
	}

	// catch Pokemon logic
	initialMsg := fmt.Sprintf("Throwing a Pokeball at %s...", pokemon)
	fmt.Println(initialMsg)

	catchChance := int((500 - pokemonResp.BaseExperience) / 5)
	randomPercent := rand.Intn(100)

	if randomPercent <= catchChance {
		caughtMsg := fmt.Sprintf("%s was caught!", pokemon)
		fmt.Println(caughtMsg)
		cfg.caughtPokemon[pokemon] = pokemonResp
		fmt.Println(cfg.caughtPokemon)
	} else {
		caughtMsg := fmt.Sprintf("%s escaped!", pokemon)
		fmt.Println(caughtMsg)
	}
	return nil
}
