package main

import "fmt"


func commandPokedex(cfg *config, param ...string) error {
	if len(cfg.pokeDex) == 0 {
		fmt.Println("Your dex is empty! Catch pokemons using the catch command.")
		return nil
	}
	for pokemonName := range cfg.pokeDex {
		fmt.Printf("  - %s\n", pokemonName)
	}
	return nil
}
