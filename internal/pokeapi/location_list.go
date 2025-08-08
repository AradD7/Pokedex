package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

)



func (c *Clinet) ListLocation(url *string) (Locations, error) {
	fullurl := baseURL + "/location-area" 
	if url != nil {
		fullurl = *url
	}
	var locations Locations
	if data, ok := c.memory.Get(fullurl); ok {
		if err := json.Unmarshal(data, &locations); err != nil {
			return Locations{}, fmt.Errorf("Could not find results in JSON: %v", err)
		}
		return locations, nil
	}
	req, err := http.NewRequest("GET", fullurl, nil)
	if err != nil{
		return Locations{}, fmt.Errorf("Couldn't make the location request: %v", err)
	}

	res, err := c.httpClient.Do(req)
	if err != nil{
		return Locations{}, fmt.Errorf("Couldn't make the location request: %v", err)
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Locations{}, fmt.Errorf("Could not read the data: %v", err)
	}

	c.memory.Add(fullurl, data)

	if err := json.Unmarshal(data, &locations); err != nil {
		return Locations{}, fmt.Errorf("Could not find results in JSON: %v", err)
	}

	return locations, nil
}


