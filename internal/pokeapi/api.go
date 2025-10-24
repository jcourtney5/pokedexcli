package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/jcourtney5/pokedexcli/internal/pokecache"
)

type LocationAreasResult struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type LocationAreaResult struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int           `json:"chance"`
				ConditionValues []interface{} `json:"condition_values"`
				MaxLevel        int           `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func PokeLocationAreasGet(url string, cache *pokecache.Cache) (LocationAreasResult, error) {
	var locationAreasResult LocationAreasResult

	// check the cache first
	cacheData, ok := cache.Get(url)
	if ok {
		fmt.Printf("Cache hit: %s\n", url)

		// Unmarshal the JSON into the struct
		err := json.Unmarshal(cacheData, &locationAreasResult)
		if err != nil {
			fmt.Println("Error unmarshalling JSON:", err)
			return locationAreasResult, err
		}
	}

	// Call the pokemon location-area API with a get request
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return locationAreasResult, err
	}
	defer res.Body.Close()

	// Check for a successful status code
	if res.StatusCode != http.StatusOK {
		err := fmt.Errorf("Response failed with status code: %d\n", res.StatusCode)
		return locationAreasResult, err
	}

	// Read the response body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return locationAreasResult, err
	}

	// Add to the cache after a read
	cache.Add(url, body)

	// Unmarshal the JSON into the struct
	err = json.Unmarshal(body, &locationAreasResult)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return locationAreasResult, err
	}

	return locationAreasResult, nil
}

func PokeLocationAreaGet(areaName string, cache *pokecache.Cache) (LocationAreaResult, error) {
	url := "https://pokeapi.co/api/v2/location-area/" + areaName

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

	// Call the pokemon API with a get request
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
