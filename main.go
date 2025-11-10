package main

import (
	"time"

	pokecache "github.com/Rhyster42/Pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokecache.NewClient(5*time.Second, time.Minute*5)
	cfg := &config{
		pokeapiClient: pokeClient,
	}

	startRepl(cfg)

}
