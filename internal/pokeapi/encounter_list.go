package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocation(locationName string) (Location, error) {
	url := baseURL + "/location-area/" + locationName

	if val, ok := c.cache.Get(url); ok {
		localLocationEncounters := Location{}
		err := json.Unmarshal(val, &localLocationEncounters)
		if err != nil {
			return Location{}, err
		}
		return localLocationEncounters, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Location{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Location{}, err
	}

	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Location{}, err
	}

	localLocationEncounters := Location{}
	err = json.Unmarshal(dat, &localLocationEncounters)
	if err != nil {
		return localLocationEncounters, err
	}

	c.cache.Add(url, dat)

	return localLocationEncounters, nil
}
