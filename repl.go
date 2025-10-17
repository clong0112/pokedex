package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(s string) []string {
	return strings.Fields(strings.ToLower(s))
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin) // implements io.Reader interface
	for {
		fmt.Print("Enter text >>")

		scanner.Scan()
		text := scanner.Text()
		cleaned := cleanInput(text)

		if len(cleaned) == 0 {
			continue
		}
		commandName := cleaned[0]
		args := []string{}
		if len(cleaned) > 1 {
			args = cleaned[1:]
		}

		availableCommands := getCommands()

		command, ok := availableCommands[commandName]
		if !ok {
			fmt.Println("Invalid command")
			continue
		}
		err := command.callback(cfg, args...)
		if err != nil {
			fmt.Println(err)
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Prints the help menu",
			callback:    callbackHelp,
		},

		"map": {
			name:        "map",
			description: "Shows the next 20 location areas",
			callback:    callbackMap,
		},

		"mapb": {
			name:        "mapb",
			description: "Shows the previous 20 location areas, if any",
			callback:    callbackMapb,
		},

		"explore": {
			name:        "explore {location_area}",
			description: "List the pokemon in a location area",
			callback:    callbackExplore,
		},

		"catch": {
			name:        "catch {pokemon_name}",
			description: "Attempt to catch a pokemon and add it to your pokedex",
			callback:    callbackCatch,
		},

		"inspect": {
			name:        "inspect {pokemon_name}",
			description: "More information about a pokemon in the pokedex",
			callback:    callbackInspect,
		},

		"pokedex": {
			name:        "pokedex",
			description: "View all pokemons in the pokedex",
			callback:    callbackPokedex,
		},

		"exit": {
			name:        "exit",
			description: "Quit the program",
			callback:    callbackExit,
		},
	}
}
