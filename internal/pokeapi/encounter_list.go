package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)



func (c *Clinet) ListEncounters(locationArea string) (Encounters, error) {
	url := baseURL + "/location-area/" + locationArea

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Encounters{}, fmt.Errorf("Could not find encounter in the given location: %v", err)
	}

	var encounters Encounters

	if data, ok := c.memory.Get(url); ok {
		if err := json.Unmarshal(data, &encounters); err != nil {
			return Encounters{}, fmt.Errorf("something went wrong with unmarshalling the data: %v", err)
		}
		return encounters, nil
	}
	
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Encounters{}, fmt.Errorf("Request to get encounters failed: %v", err)
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Encounters{}, fmt.Errorf("Reading response data failed: %v", err)
	}

	if err := json.Unmarshal(data, &encounters); err != nil {
		return Encounters{}, fmt.Errorf("something went wrong with unmarshalling the data: %v", err)
	}

	c.memory.Add(url, data)

	return encounters, nil
}

