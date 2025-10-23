package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	conf := config{
		next:     "https://pokeapi.co/api/v2/location-area",
		previous: "",
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
						"Command: %s failed with error: %w\n",
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
