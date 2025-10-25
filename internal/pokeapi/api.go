package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/jcourtney5/pokedexcli/internal/pokecache"
)

func GetLocationAreas(url string, cache *pokecache.Cache) (LocationAreas, error) {
	var locationAreas LocationAreas

	// check the cache first
	cacheData, ok := cache.Get(url)
	if ok {
		fmt.Printf("Cache hit: %s\n", url)

		// Unmarshal the JSON into the struct
		err := json.Unmarshal(cacheData, &locationAreas)
		if err != nil {
			fmt.Println("Error unmarshalling JSON:", err)
			return locationAreas, err
		}
	}

	// Call the pokemon location-area API with a get request
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return locationAreas, err
	}
	defer res.Body.Close()

	// Check for a successful status code
	if res.StatusCode != http.StatusOK {
		err := fmt.Errorf("Response failed with status code: %d\n", res.StatusCode)
		return locationAreas, err
	}

	// Read the response body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return locationAreas, err
	}

	// Add to the cache after a read
	cache.Add(url, body)

	// Unmarshal the JSON into the struct
	err = json.Unmarshal(body, &locationAreas)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return locationAreas, err
	}

	return locationAreas, nil
}

func GetLocationArea(areaName string, cache *pokecache.Cache) (LocationArea, error) {
	url := "https://pokeapi.co/api/v2/location-area/" + areaName

	var locationArea LocationArea

	// check the cache first
	cacheData, ok := cache.Get(url)
	if ok {
		fmt.Printf("Cache hit: %s\n", url)

		// Unmarshal the JSON into the struct
		err := json.Unmarshal(cacheData, &locationArea)
		if err != nil {
			fmt.Println("Error unmarshalling JSON:", err)
			return locationArea, err
		}
	}

	// Call the pokemon API with a get request
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return locationArea, err
	}
	defer res.Body.Close()

	// Check for a successful status code
	if res.StatusCode != http.StatusOK {
		err := fmt.Errorf("Response failed with status code: %d\n", res.StatusCode)
		return locationArea, err
	}

	// Read the response body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return locationArea, err
	}

	// Add to the cache after a read
	cache.Add(url, body)

	// Unmarshal the JSON into the struct
	err = json.Unmarshal(body, &locationArea)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return locationArea, err
	}

	return locationArea, nil
}

func GetPokemon(name string, cache *pokecache.Cache) (Pokemon, error) {
	url := "https://pokeapi.co/api/v2/pokemon/" + name

	var pokemon Pokemon

	// check the cache first
	cacheData, ok := cache.Get(url)
	if ok {
		fmt.Printf("Cache hit: %s\n", url)

		// Unmarshal the JSON into the struct
		err := json.Unmarshal(cacheData, &pokemon)
		if err != nil {
			fmt.Println("Error unmarshalling JSON:", err)
			return pokemon, err
		}
	}

	// Call the pokemon API with a get request
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return pokemon, err
	}
	defer res.Body.Close()

	// Check for a successful status code
	if res.StatusCode != http.StatusOK {
		err := fmt.Errorf("Response failed with status code: %d\n", res.StatusCode)
		return pokemon, err
	}

	// Read the response body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return pokemon, err
	}

	// Add to the cache after a read
	cache.Add(url, body)

	// Unmarshal the JSON into the struct
	err = json.Unmarshal(body, &pokemon)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return pokemon, err
	}

	return pokemon, nil
}
