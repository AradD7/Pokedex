package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)


type cliCommand struct {
	name			string
	description		string
	callback		func(url *config) error
}

type config struct {
	Next		string	`json:"next"`
	Previous	string	`json:"previous"`
}

var urls = config{
	Next: "https://pokeapi.co/api/v2/location-area/",
	Previous: "",
}


func cleanInput(text string) []string {
	seperatedText := strings.Split(strings.Trim(strings.ToLower(text), " "), " ")
	seperatedText = slices.DeleteFunc(seperatedText, func(s string) bool {
		return s == ""
	})
	return seperatedText
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("Pokedex > ")

		scanner.Scan()

		userInputSlice := cleanInput(scanner.Text())
		if len(userInputSlice) == 0 {
			continue
		}

		command := userInputSlice[0]
		if c, ok := getCommands()[command]; ok {
			fmt.Println()
			c.callback(&urls)
			fmt.Println()
		} else {
			fmt.Println("Unknown Command")
			}
	}
}


func getCommands() map[string]cliCommand {
	return map[string]cliCommand {
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
		"help": {
			name: "help",
			description: "Displays help message",
			callback: commandHelp,
		},
		"map": {
			name: "map",
			description: "Displays the next 20 location areas",
			callback: commandMap,
		},
		"mapb": {
			name: "mapb",
			description: "Displays the previous 20 location areas",
			callback: commandMapb,
		},
	}
}
 
