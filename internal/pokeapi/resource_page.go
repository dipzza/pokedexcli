package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

type ResourcePage struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

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
