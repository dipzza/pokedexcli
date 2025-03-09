package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/dipzza/pokedexcli/internal/pokecache"
)

const baseURL = "https://pokeapi.co/api/v2/"

const LocationEndpoint = baseURL + "location/"
const LocationAreaEndpoint = baseURL + "location-area/"

type ResourcePage struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

var cache = pokecache.NewCache(time.Second * 5)

func GetResourcePage(endpoint string) (ResourcePage, error) {
	var resource ResourcePage

	entry, found := cache.Get(endpoint)
	if !found {
		res, err := http.Get(endpoint)
		if err != nil {
			return resource, err
		}
		defer res.Body.Close()

		entry, err = io.ReadAll(res.Body)
		if err != nil {
			return resource, err
		}
		cache.Add(endpoint, entry)
	}

	if err := json.Unmarshal(entry, &resource); err != nil {
		return resource, err
	}

	return resource, nil
}