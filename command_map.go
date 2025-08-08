package main

import (
	"fmt"

)

func commandMapf(cfg *config) error {
	locationResp, err := cfg.pokeapiClient.ListLocation(cfg.nextLocationURL)
	if err != nil {
		return err
	}

	cfg.nextLocationURL = locationResp.Next
	cfg.previousLocationURL = locationResp.Previous

	for _, location := range locationResp.Results {
		fmt.Println(location.Name)
	}

	return nil

}



func commandMapb(cfg *config) error {
	if cfg.previousLocationURL == nil {
		return fmt.Errorf("You are on the first page!")
	}

	
	locationResp, err := cfg.pokeapiClient.ListLocation(cfg.previousLocationURL)
	if err != nil {
		return err
	}

	cfg.nextLocationURL = locationResp.Next
	cfg.previousLocationURL = locationResp.Previous

	for _, location := range locationResp.Results {
		fmt.Println(location.Name)
	}

	return nil

}
