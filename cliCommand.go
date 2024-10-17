package main

import (
	"fmt"
	"internal/location"
	"internal/locationArea"
	"internal/pokedex"
	"internal/pokemon"
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

func commandCatch(args string) error {
	args_list := strings.Split(args, " ")
	pokemonName := args_list[0]

	current, err := locationarea.GetCurrentLocationArea()
	if err != nil {
		fmt.Println(err)
		return err
	}

	for _, p := range current.PokemonEncounters {
		if p.Pokemon.Name == pokemonName {
			poke, err := pokemon.GetPokemon(pokemonName)
			if err != nil {
				fmt.Println(err)
				return err
			}
			fmt.Printf("Throwing a pokeball at %s...\n", poke.Name)
			if pokemon.AttemptCatch(poke) {
				fmt.Printf("Success! Caught %s\n", poke.Name)
				pokedex.Add(poke)
				return nil
			}
			fmt.Printf("%s escaped!\n", poke.Name)
			return nil
		}
	}
	fmt.Println("Cant find that pokemon in this area")
	return nil
}

func commandAreas(args string) error {
	args_list := strings.Split(args, " ")
	l := args_list[0]

	location, err := location.GetLocation(l)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Printf("%s Areas:\n", location.Name)
	for _, a := range location.Areas {
		fmt.Printf(" - %s\n", a.Name)
	}

	return nil
}

func commandInspect(args string) error {
	args_list := strings.Split(args, " ")
	pokemonName := args_list[0]

	poke, err := pokedex.Get(pokemonName)
	if err != nil {
		fmt.Println("You have not caught that pokemon")
		return err
	}
	fmt.Printf("Name: %s\n", poke.Name)
	fmt.Printf("Height: %d\n", poke.Height)
	fmt.Printf("Weight: %d\n", poke.Weight)
	fmt.Println("Stats:")
	for _, s := range poke.Stats {
		fmt.Printf(" - %s: %d (+%d)\n", s.Stat.Name, s.BaseStat, s.Effort)
	}
	fmt.Println("Types:")
	for _, t := range poke.Types {
		fmt.Printf(" - %s\n", t.Type.Name)
	}
	return nil
}

func commandPokedex(args string) error {
	fmt.Println("Your Pokedex:")
	for _, p := range pokedex.GetPokedex() {
		fmt.Printf(" - %s\n", p.Name)
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
		"catch": {
			name:        "catch",
			description: "Attempts to catch a pokemon and then adds its info to the pokedex",
			callback:    commandCatch,
			args: map[string]cliCommand{
				"test": {
					name:        "test",
					description: "test",
					callback:    commandExit,
					args:        nil,
				},
			},
		},
		"areas": {
			name:        "areas",
			description: "Displays the areas in a specific location",
			callback:    commandAreas,
			args: map[string]cliCommand{
				"test": {
					name:        "test",
					description: "test",
					callback:    commandExit,
					args:        nil,
				},
			},
		},
		"inspect": {
			name:        "inspect",
			description: "Displays information about pokemon in pokedex",
			callback:    commandInspect,
			args: map[string]cliCommand{
				"test": {
					name:        "test",
					description: "test",
					callback:    commandExit,
					args:        nil,
				},
			},
		},
		"pokedex": {
			name:        "pokedex",
			description: "Displays all of the pokemon in your pokedex",
			callback:    commandPokedex,
			args: map[string]cliCommand{
				"test": {
					name:        "test",
					description: "test",
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
