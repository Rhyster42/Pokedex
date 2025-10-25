package main

import (
	"fmt"

	requests "github.com/Rhyster42/Pokedex/internal"
)

func commandMap(cfg *config) error {

	var url string

	if cfg.Next == "" {
		url = "https://pokeapi.co/api/v2/location-area/"
	} else {
		url = cfg.Next
	}

	data, err := requests.GetData(url)
	if err != nil {
		return err
	}

	for i := 0; i < len(data.Results); i++ {
		fmt.Println(data.Results[i].Name)
	}

	cfg.Next = data.Next
	cfg.Previous = data.Previous

	return nil
}

func commandMapb(cfg *config) error {
	if cfg.Previous == "" {
		fmt.Println("No Previous Locations")
		return nil
	}
	cfg.Next = cfg.Previous
	commandMap(cfg)
	return nil
}
