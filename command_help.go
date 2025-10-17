package main

import "fmt"

func callbackHelp(cfg *config, args ...string) error {
	fmt.Println("Welcome to Pokedex: ")
	fmt.Println("List of available commands: ")
	for _, cmd := range(getCommands()) {
		fmt.Println("=> ", cmd.name, ": ", cmd.description)
	}
	return nil
}