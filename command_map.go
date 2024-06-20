package main

import "github.com/alexandreffaria/pokedex/pokeapi"

func commandMap() error {
	pokeapi.FetchLocations()
	return nil
}
