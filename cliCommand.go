package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var apiURL = "https://pokeapi.co/api/v2/"
var currentLocation = "location/"

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
	locations, err := getLocations(currentLocation)
	if err != nil {
		return err
	}
	for _, l := range locations {
		fmt.Println(l.Name)
	}
	return nil
}

func getLocation

func getLocations(direction string) ([]Location, error) {
	fullURL := apiURL + currentLocation

	if direction == "previous" {

	}

	// Create a new get request at the location endpoint
	res, err := http.Get(fullURL)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println(res.Status)
		return nil, fmt.Errorf(res.Status)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var response Response
	if err = json.Unmarshal(body, &response); err != nil {
		fmt.Println(err)
		return nil, err
	}

	//fmt.Println(response)
	return response.Results, nil
}

func commandMapb() error {
	locations, err := getLocations()
	if err != nil {
		return err
	}
	for _, l := range locations {
		fmt.Println(l.Name)
	}
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
