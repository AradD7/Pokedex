package main

import (
	"fmt"
	"time"
)


func commandExplore(cfg *config, param ...string) error {
	if len(param) != 2 {
		return fmt.Errorf("Please enter a location to explore")
	}
	encounters, err := cfg.pokeapiClient.ListEncounters(param[1])
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s", param[1])
	time.Sleep(time.Second)
	fmt.Printf(".")
	time.Sleep(time.Second)
	fmt.Printf(".")
	time.Sleep(time.Second)
	fmt.Printf(".")
	time.Sleep(time.Second)
	fmt.Printf("\nFound Pokemon: \n")


	for _, encounter := range encounters.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	} 
	return nil
}
