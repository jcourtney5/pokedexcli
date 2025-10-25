package main

import (
	"fmt"
	"github.com/jcourtney5/pokedexcli/internal/pokeapi"
	"github.com/jcourtney5/pokedexcli/internal/pokecache"
	"math/rand"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(conf *config, arg1 string) error
}

type config struct {
	next     string
	previous string
	cache    *pokecache.Cache
	pokedex  map[string]pokeapi.Pokemon
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
	"explore": {
		name:        "explore",
		description: "Find a list of all pokemon at the given location <area_name>",
		callback:    commandExplore,
	},
	"catch": {
		name:        "catch",
		description: "Attempt to catch pokemon with the given <pokemon_name>",
		callback:    commandCatch,
	},
}

func getCommand(name string) (cliCommand, bool) {
	value, ok := commands[name]
	return value, ok
}

func commandExit(conf *config, arg1 string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(conf *config, arg1 string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	fmt.Println("help: Displays a help message")
	fmt.Println("exit: Exit the Pokedex")
	fmt.Println("map: Displays the names of next 20 location areas in the Pokemon world")
	fmt.Println("mapb: Displays the names of previous 20 location areas in the Pokemon world")
	fmt.Println("explore <area_name>: Find a list of all pokemon at the given location <area_name>")
	fmt.Println("catch <pokemon_name>: Attempt to catch pokemon with the given <pokemon_name>")
	return nil
}

func commandMap(conf *config, arg1 string) error {
	url := conf.next
	if len(conf.next) == 0 {
		url = "https://pokeapi.co/api/v2/location-area"
	}

	// call the poke API with next
	locationAreas, err := pokeapi.GetLocationAreas(url, conf.cache)
	if err != nil {
		return fmt.Errorf("Error calling poke pokeapi %s, failed with error : %w\n", conf.next, err)
	}

	// print the results
	for i := 0; i < len(locationAreas.Results); i++ {
		fmt.Println(locationAreas.Results[i].Name)
	}

	// save new previous and next
	if locationAreas.Previous != nil {
		conf.previous = *locationAreas.Previous
	} else {
		conf.previous = ""
	}
	if locationAreas.Next != nil {
		conf.next = *locationAreas.Next
	} else {
		conf.next = ""
	}

	return nil
}

func commandMapB(conf *config, arg1 string) error {
	if len(conf.previous) == 0 {
		fmt.Println("you're on the first page")
		return nil
	} else {
		// call the poke API with previous
		res, err := pokeapi.GetLocationAreas(conf.previous, conf.cache)
		if err != nil {
			return fmt.Errorf("Error calling pokeapi.GetLocationArea, failed with error : %w\n", err)
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

func commandExplore(conf *config, arg1 string) error {
	fmt.Printf("Exploring %s...\n", arg1)

	// call the location area api with the area name
	locationArea, err := pokeapi.GetLocationArea(arg1, conf.cache)
	if err != nil {
		return fmt.Errorf("Error calling pokeapi.GetLocationArea, failed with error : %w\n", err)
	}

	// print the results
	fmt.Println("Found Pokemon:")
	for i := 0; i < len(locationArea.PokemonEncounters); i++ {
		fmt.Printf(" - %s\n", locationArea.PokemonEncounters[i].Pokemon.Name)
	}

	return nil
}

func commandCatch(conf *config, arg1 string) error {
	fmt.Printf("Throwing a Pokeball at %s...\n", arg1)

	// call the pokemon get API to get details about the pokemon
	pokemon, err := pokeapi.GetPokemon(arg1, conf.cache)
	if err != nil {
		return fmt.Errorf("Error calling pokeapi.GetPokemon, failed with error : %w\n", err)
	}

	// User base experience and a random number between 10 and 255 to try and catch
	base_experience := pokemon.BaseExperience

	// Generate a random integer between 0 and 200 (inclusive)
	rand_num := rand.Intn(200)

	fmt.Printf("base_experience: %v, your roll: %v\n", base_experience, rand_num)

	// See if we caught the pokemon
	if rand_num > base_experience {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		conf.pokedex[pokemon.Name] = pokemon

		fmt.Println("Current caught pokemon:")
		for name, _ := range conf.pokedex {
			fmt.Printf(" - %s\n", name)
		}
	} else {
		fmt.Printf("%s was not caught!\n", pokemon.Name)
	}

	return nil
}
