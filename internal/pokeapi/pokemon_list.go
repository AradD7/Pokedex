package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)



func (c *Clinet) PokemonList(pokemonName string) (PokemonType, error) {
	url := baseURL + "/pokemon/" + pokemonName

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonType{}, fmt.Errorf("Couldn't create a request to find pokemon: %v", err)
	}

	var pokemon PokemonType

	if data, ok := c.memory.Get(url); ok {
		if err := json.Unmarshal(data, &pokemon); err != nil {
			return PokemonType{}, fmt.Errorf("something went wrong with unmarshalling the data: %v", err)
		}
		return pokemon, nil
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonType{}, fmt.Errorf("Get request for pokemon failed: %v", err)
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonType{}, fmt.Errorf("Could not read the data returned for pokemon: %v", err)
	}
	
	if err := json.Unmarshal(data, &pokemon); err != nil {
		return PokemonType{}, fmt.Errorf("something went wrong with unmarshalling the data: %v", err)
	}
	c.memory.Add(url, data)

	return pokemon, nil


}
