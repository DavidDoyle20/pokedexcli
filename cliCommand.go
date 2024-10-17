package main

import (
	"fmt"
	"internal/location"
	"internal/locationArea"
	"os"
	"strings"
)

type CommandStack struct {
	args []cliCommand
}

type cliCommand struct {
	name        string
	description string
	args        map[string]cliCommand
	callback    func(string) error
}

func commandHelp(args string) error {
	fmt.Println("Usage: ")
	fmt.Println()
	for _, v := range getCommandMap() {
		fmt.Printf("%s: %s", v.name, v.description)
		fmt.Println()
	}
	fmt.Println()

	return nil
}

func commandExit(args string) error {
	os.Exit(0)
	return nil
}

func commandMap(args string) error {
	resp, err := location.GetLocations()
	if err != nil {
		return err
	}
	for _, l := range resp {
		fmt.Println(l.Name)
	}

	return nil
}

func commandMapb(args string) error {
	resp, err := location.GetPreviousLocations()
	if err != nil {
		return err
	}

	for _, l := range resp {
		fmt.Println(l.Name)
	}

	return nil
}

func commandExplore(args string) error {
	args_list := strings.Split(args, " ")
	area := args_list[0]

	la, err := locationarea.GetLocationArea(area)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", la.Name)

	err = commandNames(la)
	if err != nil {
		return err
	}
	return nil
}

func commandNames(la locationarea.LocationArea) error {
	fmt.Println("Found Pokemon:")
	for _, p := range la.PokemonEncounters {
		fmt.Printf(" - %s\n", p.Pokemon.Name)
	}
	return nil
}

func getCommandMap() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
			args: map[string]cliCommand{
				"test": {
					name:        "test",
					description: "test",
					callback:    commandExit,
					args:        nil,
				},
			},
		},
		"exit": {
			name:        "exit",
			description: "Exits the pokedex",
			callback:    commandExit,
			args: map[string]cliCommand{
				"test": {
					name:        "test",
					description: "test",
					callback:    commandExit,
					args:        nil,
				},
			},
		},
		"map": {
			name:        "map",
			description: "Displays the names of the next 20 locations in the pokemon world",
			callback:    commandMap,
			args: map[string]cliCommand{
				"test": {
					name:        "test",
					description: "test",
					callback:    commandExit,
					args:        nil,
				},
			},
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 locations in the pokemon world",
			callback:    commandMapb,
			args: map[string]cliCommand{
				"test": {
					name:        "test",
					description: "test",
					callback:    commandExit,
					args:        nil,
				},
			},
		},
		"explore": {
			name:        "explore",
			description: "Display the pokemon that can be found in an area",
			callback:    commandExplore,
			args: map[string]cliCommand{
				"-n": {
					name:        "names",
					description: "Lists the names of the pokemon that can be encountered at a location",
					callback:    commandExit,
					args:        nil,
				},
			},
		},
	}
}

// The main method should call this method on user input
// If the command is in valid commands return it
func getCommand(s string) (cliCommand, error) {
	fields := strings.SplitAfterN(s, " ", 2)
	c := strings.TrimSpace(fields[0])
	args := ""
	if len(fields) > 1 {
		args = strings.TrimRight(fields[1], " ")
	}

	cmd, OK := getCommandMap()[c]
	if !OK {
		return cliCommand{}, fmt.Errorf("Command not found: %s", s)
	}

	cmd.callback(args)
	return cmd, nil
}
