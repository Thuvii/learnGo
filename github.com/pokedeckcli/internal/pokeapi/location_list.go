package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocation(pageURL *string) (LocationStruct, error) {
	locationResp := LocationStruct{}

	url := baseURL + "location-area"
	if pageURL != nil {
		url = *pageURL
	}

	val, ok := c.cache.Get(url)
	if ok {
		if err := json.Unmarshal(val, &locationResp); err != nil {
			return LocationStruct{}, err
		}
		return locationResp, nil
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationStruct{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationStruct{}, err
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationStruct{}, err
	}

	c.cache.Add(url, data)

	if err := json.Unmarshal(data, &locationResp); err != nil {
		return LocationStruct{}, err
	}

	return locationResp, nil
}
