package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"errors"
	"io"
	"net/http"
	"encoding/json"
)

var ErrExit = errors.New("exit")

var commands map[string]cliCommand

func init() {
	commands = make(map[string]cliCommand)

	commands["help"] = cliCommand{
		name: "help",
		description: "Displays a help message",
		callback: commandHelp,
	}

	commands["exit"] = cliCommand{
		name:        "exit",
		description: "Exit the Pokedex",
		callback: commandExit,
	}

	commands["map"] = cliCommand{
		name: "map",
		description: "Go forward to the next 20 areas",
		callback: commandMap,
	}
	
	commands["mapb"] = cliCommand{
		name: "mapb",
		description: "Go back to the previous 20 areas",
		callback: commandMapb,
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	c := &Config{}

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		cleanedInput := cleanInput(scanner.Text())
		if len(cleanedInput) == 0 { continue }

		cmd, ok := commands[cleanedInput[0]]

		if !ok { fmt.Println("Unknown command"); continue }

		if err := cmd.callback(c); err != nil { fmt.Println("Error:", err) }
	}
}

func cleanInput (text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func commandMapb(c *Config) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if c.Previous != "" {
		url = c.Previous
	} else {
		fmt.Println("You're on the first page!")
	}

	resp, err := GetLocationArea(url)
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

func commandMap(c *Config) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if c.Next != "" {
		url = c.Next
	}

	resp, err := GetLocationArea(url)
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

func commandExit(c *Config) error {
    fmt.Println("Closing the Pokedex... Goodbye!")
    os.Exit(0)
    return nil
}


func commandHelp(c *Config) error {
	fmt.Printf(`Welcome to the Pokedex!
Usage:

`)
	for k, v := range commands {
		fmt.Printf("%v: %v\n", k, v.description)
	}
	return nil
}
	
type cliCommand struct {
	name string
	description string
	callback func(*Config) error
}

type Config struct {
	Next string
	Previous string
}

type LocationAreas struct {
		Next *string 
		Previous *string			
		Results []NameUrl
}

type NameUrl struct {
		Name string
		URL string
}

func GetLocationArea(url string) (LocationAreas, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreas{}, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return LocationAreas{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreas{}, err
	}

	var data LocationAreas
	err = json.Unmarshal(body, &data)
	if err != nil {
		return LocationAreas{}, err
	}

	return data, nil
}
