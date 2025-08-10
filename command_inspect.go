package main

import (
	"fmt"

	"github.com/AradD7/Pokedex/internal/pokeapi"
)


func commandInspect(cfg *config, param ...string) error {
	if len(param) != 2 {
		return fmt.Errorf("Please enter the pokemon name")
	}
	pokemonName := param[1]

	if pokemon, ok := cfg.pokeDex[pokemonName]; ok {
		printPokemonStats(pokemon)
	} else {
		fmt.Println("You have not caught", pokemonName)
	}
	return nil
}


func printPokemonStats(pokemon pokeapi.PokemonType) {
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Printf("Stats:\n")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Printf("Types:\n")
	for _, typeOfPokemon := range pokemon.Types {
		fmt.Printf("  - %s\n", typeOfPokemon.Type.Name)
	}
}
