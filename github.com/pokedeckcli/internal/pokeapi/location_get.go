package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListPokebyArea(area string) (EncounterStruct, error) {
	url := baseURL + "location-area/" + area
	pokeList := EncounterStruct{}

	if val, ok := c.cache.Get(url); ok {
		if err := json.Unmarshal(val, &pokeList); err != nil {
			return EncounterStruct{}, err
		}
		return pokeList, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return EncounterStruct{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return EncounterStruct{}, err
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return EncounterStruct{}, err
	}
	c.cache.Add(url, data)

	if err := json.Unmarshal(data, &pokeList); err != nil {
		return EncounterStruct{}, err
	}

	return pokeList, nil

}
