package main

import (
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func commandHelp() error {

	fmt.Println("Usage: \n")
	for _, v := range getCommandMap() {
		fmt.Printf("%s: %s", v.name, v.description)
		fmt.Println()
	}
	fmt.Println()

	return nil
}

func commandExit() error {
	os.Exit(0)
	return nil
}

func commandMap() error {
	resp, err := getResponse((nextLocation))
	if err != nil {
		return err
	}
	for _, l := range resp.Results {
		fmt.Println(l.Name)
	}

	currentLocation = nextLocation
	nextLocation = resp.Next
	previousLocation = resp.Previous

	return nil
}

func commandMapb() error {
	if previousLocation == "" {
		nextLocation = currentLocation
		err := fmt.Errorf("No previous found")
		fmt.Println(err)
		return err
	}

	resp, err := getResponse((previousLocation))
	if err != nil {
		return err
	}

	for _, l := range resp.Results {
		fmt.Println(l.Name)
	}

	currentLocation = previousLocation
	nextLocation = resp.Next
	previousLocation = resp.Previous

	return nil
}

func getCommandMap() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exits the pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays the names of the next 20 locations in the pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 locations in the pokemon world",
			callback:    commandMapb,
		},
	}
}

// The main method should call this method on user input
// If the command is in valid commands return it
func getCommand(s string) (cliCommand, error) {
	s = strings.Fields(s)[0]
	cmd, OK := getCommandMap()[s]
	if !OK {
		return cliCommand{}, fmt.Errorf("Command not found: %s", s)
	}
	return cmd, nil
}
