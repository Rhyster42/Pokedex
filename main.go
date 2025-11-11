package main

import (
	"time"

	"github.com/Rhyster42/Pokedex/internal/pokeapi"
	pokecache "github.com/Rhyster42/Pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokecache.NewClient(5*time.Second, time.Minute*5)
	cfg := &config{
		caughtPokemon: map[string]pokeapi.PokemonInfo{},
		pokeapiClient: pokeClient,
	}

	startRepl(cfg)

}
