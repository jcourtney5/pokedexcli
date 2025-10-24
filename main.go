package main

import (
	"bufio"
	"fmt"
	"github.com/jcourtney5/pokedexcli/internal/pokecache"
	"os"
	"time"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	conf := config{
		next:     "https://pokeapi.co/api/v2/location-area",
		previous: "",
		cache:    pokecache.NewCache(15 * time.Second),
	}

	for {
		fmt.Print("Pokedex > ")

		if !scanner.Scan() {
			break
		}

		input := scanner.Text()
		words := cleanInput(input)
		if len(words) > 0 {
			command := words[0]
			cliCommand, ok := getCommand(command)
			if ok {
				err := cliCommand.callback(&conf)
				if err != nil {
					fmt.Printf(
						"Command: %s failed with error: %v\n",
						command,
						err,
					)
				}
			} else {
				fmt.Printf("Unknown command: %s\n", command)
			}
		}
	}

	// Check for any error that might’ve occurred
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
	}
}
