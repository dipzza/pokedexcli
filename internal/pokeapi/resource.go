package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func getResource[T any](endpoint string) (T, error) {
	var resource T

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