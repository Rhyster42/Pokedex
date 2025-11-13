package main

import "fmt"

func commandPokedex(cfg *config, _ string) error {
	if cfg.caughtPokemon == nil {
		fmt.Println("You're Pokedex is empty. Go catch some Pokemon!")
		return nil
	}
	for _, pokemon := range cfg.caughtPokemon {
		fmt.Printf("   - %s\n", pokemon.Name)
	}
	return nil
}
