package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type locationResults struct {
	LocationsSlice 	[]struct {
		Name 		string `json:"name"`
		Location 	string `json:"url"`
	} `json:"results"`
}


func commandMap(urls *config) error {
	res, err := http.Get(urls.Next)
	if err != nil{
		return fmt.Errorf("Couldn't make the location request: %v", err)
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("Could not read the data: %v", err)
	}

	var locations locationResults
	if err := json.Unmarshal(data, &locations); err != nil {
		return fmt.Errorf("Could not find results in JSON: %v", err)
	}
	if err := json.Unmarshal(data, &urls); err != nil {
		return fmt.Errorf("Could not find results in JSON: %v", err)
	}

	for _, location := range locations.LocationsSlice {
		fmt.Println(location.Name)
	}

	return nil
}



func commandMapb(urls *config) error {
	res, err := http.Get(urls.Previous)
	if err != nil{
		return fmt.Errorf("Couldn't make the location request: %v", err)
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("Could not read the data: %v", err)
	}

	var locations locationResults
	if err := json.Unmarshal(data, &locations); err != nil {
		return fmt.Errorf("Could not find results in JSON: %v", err)
	}
	if err := json.Unmarshal(data, &urls); err != nil {
		return fmt.Errorf("Could not find results in JSON: %v", err)
	}

	for _, location := range locations.LocationsSlice {
		fmt.Println(location.Name)
	}

	return nil
}
