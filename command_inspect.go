package main

import (
	"fmt"
)

func commandInspect(cfg *config, pokemon string) error {
	pokeData, ok := cfg.caughtPokemon[pokemon]
	if !ok {
		fmt.Println("Pokemon not in Pokedex.")
		return nil
	}

	fmt.Printf("Name: %s\n", pokeData.Name)
	fmt.Printf("Height: %d\n", pokeData.Height)
	fmt.Printf("Weight: %d\n", pokeData.Weight)

	fmt.Println("Stats:")
	for _, val := range pokeData.Stats {
		fmt.Printf("   -%s: %d\n", val.Stat.Name, val.BaseStat)
	}

	fmt.Println("Types:")
	for _, pokeType := range pokeData.Types {
		fmt.Printf("   -%s\n", pokeType.Type.Name)
	}

	return nil
}
