package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) getPokemon(pokemon string) (PokemonInfo, error) {

	url := baseURL + "/pokemon/" + pokemon

	// cache check
	if val, ok := c.cache.Get(url); ok {
		pokemonInfo := PokemonInfo{}
		err := json.Unmarshal(val, &pokemonInfo)
		if err != nil {
			return PokemonInfo{}, err
		}
		return pokemonInfo, nil
	}

	// request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonInfo{}, err
	}

	// pull response data from request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonInfo{}, err
	}
	defer resp.Body.Close()

	// read data
	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonInfo{}, err
	}

	// unmarshal data from JSON
	pokemonInfo := LocationInfo{}
	err = json.Unmarshal(dat, &pokemonInfo)
	if err != nil {
		return PokemonInfo{}, err
	}

	//add data to cache, return
	c.cache.Add(url, dat)
	return PokemonInfo{}, nil
}
