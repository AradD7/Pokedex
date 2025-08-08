package main

import (
	"math"
	"math/rand"
	"fmt"
	"time"
)


func commandCatch(cfg *config, param ...string) error {
	if len(param) != 2 {
		return fmt.Errorf("Please enter a pokemon name to catch")
	}
	pokemon, err := cfg.pokeapiClient.PokemonList(param[1])
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s", param[1])
	time.Sleep(time.Second)
	fmt.Printf(".")
	time.Sleep(time.Second)
	fmt.Printf(".")
	time.Sleep(time.Second)
	fmt.Printf(".")
	time.Sleep(time.Second)
	chance := int(math.Round(float64(pokemon.BaseExperience) / 20))
	if rand.Intn(chance) == 0 {
		fmt.Printf("\n%s was cought!", param[1])
		cfg.pokeDex[param[1]] = pokemon
	} else {
		fmt.Printf("\n%s escaped", param[1])
	}
	return nil
}
