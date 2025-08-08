package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/AradD7/Pokedex/internal/pokeapi"
)


type cliCommand struct {
	name			string
	description		string
	callback		func(cfg *config) error
}

type config struct {
	pokeapiClient		pokeapi.Clinet
	nextLocationURL		*string
	previousLocationURL	*string
}



func cleanInput(text string) []string {
	seperatedText := strings.Split(strings.Trim(strings.ToLower(text), " "), " ")
	seperatedText = slices.DeleteFunc(seperatedText, func(s string) bool {
		return s == ""
	})
	return seperatedText
}

func startRepl(cfg *config) {
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
			err := c.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
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
			callback: commandMapf,
		},
		"mapb": {
			name: "mapb",
			description: "Displays the previous 20 location areas",
			callback: commandMapb,
		},
	}
}
 
