package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, name string) error {
	if name == "" {
		return errors.New("no pokemon argument, please add pokemon you are attempting to catch")
	}

	pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	// catch Pokemon logic
	initialMsg := fmt.Sprintf("Throwing a Pokeball at %s...", name)
	fmt.Println(initialMsg)

	catchChance := int((500 - pokemon.BaseExperience) / 5)
	randomPercent := rand.Intn(100)

	if randomPercent <= catchChance {
		caughtMsg := fmt.Sprintf("%s was caught!", pokemon.Name)
		fmt.Println(caughtMsg)
		cfg.caughtPokemon[pokemon.Name] = pokemon
	} else {
		caughtMsg := fmt.Sprintf("%s escaped!", pokemon.Name)
		fmt.Println(caughtMsg)
	}
	return nil
}
