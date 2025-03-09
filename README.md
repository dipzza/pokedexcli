# pokedexcli
Pokedex is a command-line REPL using PokéAPI

Example: 
```console
Pokedex > help
Welcome to the Pokedex!
Usage:

help: Displays a help message
map: Displays the next 20 locations in the map
mapb: Displays the previous 20 locations in the map
explore <area_name>: Displays Pokémons in an area
catch <pokemon_name>: Try to catch a pokemon
inspect <pokemon_name>: Displays pokemon information
pokedex: list all pokemon in the Pokedex
exit: Exit the Pokedex

Pokedex > catch pikachu
Throwing a Pokeball at pikachu...
pikachu was caught!

Pokedex > inspect pikachu
Name: pikachu
Height: 4
Weight: 60
Stats:
 -hp: 35
 -attack: 55
 -defense: 40
 -special-attack: 50
 -special-defense: 50
 -speed: 90
Types:
 - electric

```
