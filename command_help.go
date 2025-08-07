package main

import "fmt"


func commandHelp(urls *config) error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")
	for key, value := range getCommands() {
		fmt.Printf("%s: %s\n", key, value.description)
	}
	return nil
}
