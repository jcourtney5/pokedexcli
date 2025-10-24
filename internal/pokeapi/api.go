package pokeapi

import (
	"encoding/json"
	"fmt"
	"github.com/jcourtney5/pokedexcli/internal/pokecache"
	"io"
	"net/http"
)

type LocationAreaResult struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func PokeLocationAreaGet(url string, cache *pokecache.Cache) (LocationAreaResult, error) {
	var locationAreaResult LocationAreaResult

	// check the cache first
	cacheData, ok := cache.Get(url)
	if ok {
		fmt.Printf("Cache hit: %s\n", url)

		// Unmarshal the JSON into the struct
		err := json.Unmarshal(cacheData, &locationAreaResult)
		if err != nil {
			fmt.Println("Error unmarshalling JSON:", err)
			return locationAreaResult, err
		}
	}

	// Call the pokemon location-area API with a get request
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return locationAreaResult, err
	}
	defer res.Body.Close()

	// Check for a successful status code
	if res.StatusCode != http.StatusOK {
		err := fmt.Errorf("Response failed with status code: %d\n", res.StatusCode)
		return locationAreaResult, err
	}

	// Read the response body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return locationAreaResult, err
	}

	// Add to the cache after a read
	cache.Add(url, body)

	// Unmarshal the JSON into the struct
	err = json.Unmarshal(body, &locationAreaResult)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return locationAreaResult, err
	}

	return locationAreaResult, nil
}
