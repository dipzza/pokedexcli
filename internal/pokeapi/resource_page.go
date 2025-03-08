package pokeapi

import (
	"encoding/json"
	"net/http"
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

func GetResourcePage(endpoint string) (ResourcePage, error) {
	var resource ResourcePage
	res, err := http.Get(endpoint)
	if err != nil {
		return resource, err
	}
	defer res.Body.Close()

	
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&resource); err != nil {
		return resource, err
	}

	return resource, nil
}