package main

import (
	"fmt"
	"github.com/jcourtney5/pokedexcli/internal/api"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(conf *config) error
}

type config struct {
	next     string
	previous string
}

var commands = map[string]cliCommand{
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	},
	"help": {
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp,
	},
	"map": {
		name:        "map",
		description: "Displays the names of next 20 location areas in the Pokemon world",
		callback:    commandMap,
	},
	"mapb": {
		name:        "mapb",
		description: "Displays the names of previous 20 location areas in the Pokemon world",
		callback:    commandMapB,
	},
}

func getCommand(name string) (cliCommand, bool) {
	value, ok := commands[name]
	return value, ok
}

func commandExit(conf *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(conf *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:\n")
	fmt.Println("help: Displays a help message")
	fmt.Println("exit: Exit the Pokedex")
	fmt.Println("map: Displays the names of next 20 location areas in the Pokemon world")
	fmt.Println("mapb: Displays the names of previous 20 location areas in the Pokemon world")
	return nil
}

func commandMap(conf *config) error {
	url := conf.next
	if len(conf.next) == 0 {
		url = "https://pokeapi.co/api/v2/location-area"
	}

	// call the poke API with next
	res, err := api.PokeLocationAreaGet(url)
	if err != nil {
		return fmt.Errorf("Error calling poke api %s, failed with error : %w\n", conf.next, err)
	}

	// print the results
	for i := 0; i < len(res.Results); i++ {
		fmt.Println(res.Results[i].Name)
	}

	// save new previous and next
	if res.Previous != nil {
		conf.previous = *res.Previous
	} else {
		conf.previous = ""
	}
	if res.Next != nil {
		conf.next = *res.Next
	} else {
		conf.next = ""
	}

	return nil
}

func commandMapB(conf *config) error {
	if len(conf.previous) == 0 {
		fmt.Println("you're on the first page")
		return nil
	} else {
		// call the poke API with previous
		res, err := api.PokeLocationAreaGet(conf.previous)
		if err != nil {
			return fmt.Errorf("Error calling poke api %s, failed with error : %w\n", conf.next, err)
		}

		// print the results
		for i := 0; i < len(res.Results); i++ {
			fmt.Println(res.Results[i].Name)
		}

		// save new previous and next
		if res.Previous != nil {
			conf.previous = *res.Previous
		} else {
			conf.previous = ""
		}
		if res.Next != nil {
			conf.next = *res.Next
		} else {
			conf.next = ""
		}
	}
	return nil
}
