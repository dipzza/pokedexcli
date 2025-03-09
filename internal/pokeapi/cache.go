package pokeapi

import (
	"time"

	"github.com/dipzza/pokedexcli/internal/pokecache"
)

var cache = pokecache.NewCache(time.Second * 10)
