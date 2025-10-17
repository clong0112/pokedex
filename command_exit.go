package main

import (
	"fmt"
	"os"
)

func callbackExit(cfg *config, args ...string) error {
	fmt.Println("Exiting program ...")
	os.Exit(0)
	return nil
}
