package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) PokemonInfo(pokemonName string) (Pokemon, error) {
	url := baseURL + "pokemon/" + pokemonName
	pokemonInfo := Pokemon{}

	if val, ok := c.cache.Get(url); ok {
		if err := json.Unmarshal(val, &pokemonInfo); err != nil {
			return Pokemon{}, err
		}
		return pokemonInfo, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}
	c.cache.Add(url, data)
	if err := json.Unmarshal(data, &pokemonInfo); err != nil {
		return Pokemon{}, err
	}
	return pokemonInfo, nil
}
