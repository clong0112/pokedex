package main

import (
	"errors"
	"fmt"
)

func callbackMap(cfg *config, args ...string) error {
	resp, err := cfg.apiClient.ListLocationAreas(cfg.NextLocationAreaURL)
	if err != nil {
		return err
	}
	fmt.Println("Location Areas:")
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}
	cfg.NextLocationAreaURL = resp.Next
	cfg.PreviousLocationAreaURL = resp.Previous
	return nil
}

func callbackMapb(cfg *config, args ...string) error {
	if cfg.PreviousLocationAreaURL == nil {
		return errors.New("you're on the first page of locations")
	}
	resp, err := cfg.apiClient.ListLocationAreas(cfg.PreviousLocationAreaURL)
	if err != nil {
		return err
	}
	fmt.Println("Location Areas:")
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}
	cfg.NextLocationAreaURL = resp.Next
	cfg.PreviousLocationAreaURL = resp.Previous
	return nil
}
