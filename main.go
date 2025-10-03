package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/clong0112/pokedex/internal/apis"
	"github.com/clong0112/pokedex/internal/pokecache"
)

var ErrExit = errors.New("exit")

var commands map[string]apis.CliCommand

func init() {
	
	commands = make(map[string]apis.CliCommand)

	commands["help"] = apis.CliCommand{
		Name: "help",
		Description: "Displays a help message",
		Callback: commandHelp,
	}

	commands["exit"] = apis.CliCommand{
		Name:        "exit",
		Description: "Exit the Pokedex",
		Callback: commandExit,
	}

	commands["map"] = apis.CliCommand{
		Name: "map",
		Description: "Go forward to the next 20 areas",
		Callback: commandMap,
	}
	
	commands["mapb"] = apis.CliCommand{
		Name: "mapb",
		Description: "Go back to the previous 20 areas",
		Callback: commandMapb,
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	c := &apis.Config{Cache: pokecache.NewCache(time.Second * 5)}
	

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		cleanedInput := cleanInput(scanner.Text())
		if len(cleanedInput) == 0 { continue }

		cmd, ok := commands[cleanedInput[0]]

		if !ok { fmt.Println("Unknown command"); continue }

		if err := cmd.Callback(c); err != nil { fmt.Println("Error:", err) }
	}
}

func cleanInput (text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func commandMapb(c *apis.Config) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if c.Previous != "" {
		url = c.Previous
	} else {
		fmt.Println("You're on the first page!")
	}

	resp, err := apis.GetLocationArea(url, c)
	if err != nil { return err }

	for _, r := range resp.Results {
		fmt.Println(r.Name)
	}

	if resp.Next != nil {
		c.Next = *resp.Next
	} else {
		c.Next = ""
	}

	if resp.Previous != nil {
		c.Previous = *resp.Previous
	} else {
		c.Previous = ""
	}

	return nil
}

func commandMap(c *apis.Config) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if c.Next != "" {
		url = c.Next
	}

	resp, err := apis.GetLocationArea(url, c)
	if err != nil { return err }

	for _, r := range resp.Results {
		fmt.Println(r.Name)
	}

	if resp.Next != nil {
		c.Next = *resp.Next
	} else {
		c.Next = ""
	}
	
	if resp.Previous != nil {
		c.Previous = *resp.Previous
	} else {
		c.Previous = ""
	}
	return nil
}

func commandExit(c *apis.Config) error {
    fmt.Println("Closing the Pokedex... Goodbye!")
    os.Exit(0)
    return nil
}


func commandHelp(c *apis.Config) error {
	fmt.Printf(`Welcome to the Pokedex!
Usage:

`)
	for k, v := range commands {
		fmt.Printf("%v: %v\n", k, v.Description)
	}
	return nil
}