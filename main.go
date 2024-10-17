package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println()

	// Returns a new scanner that reads from stdin
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("pokedex > ")

		// First we read the user input
		if !scanner.Scan() {
			fmt.Println("Scan failed!")
			continue
		}

		if err := scanner.Err(); err != nil {
			fmt.Println(err)
			continue
		}

		_, err := getCommand(cleanInput(scanner.Text()))
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
}

func cleanInput(s string) string {
	return strings.ToLower(s)
}
