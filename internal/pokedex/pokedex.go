package pokedex

import (
	"fmt"
	"internal/pokemon"
)

var pokedex = make(map[string]pokemon.Pokemon)

func Add(poke pokemon.Pokemon) error {
	_, OK := pokedex[poke.Name]
	if !OK {
		pokedex[poke.Name] = poke
		fmt.Printf("Added %s to the pokedex!\n", poke.Name)
	}
	return nil
}

func Get(name string) (pokemon.Pokemon, error) {
	poke, OK := pokedex[name]
	if !OK {
		return pokemon.Pokemon{}, fmt.Errorf("Pokemon not in pokedex")
	}
	return poke, nil
}

func GetPokedex() map[string]pokemon.Pokemon {
	return pokedex
}
