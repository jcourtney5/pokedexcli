package main

import (
	"strings"
)

func cleanInput(text string) []string {
	result := []string{}

	words := strings.Fields(text)
	for _, word := range words {
		if len(word) > 0 {
			wordFixed := strings.ToLower(strings.TrimSpace(word))
			result = append(result, wordFixed)
		}
	}

	return result
}
