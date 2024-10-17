module github.com/DavidDoyle20/pokedexcli

go 1.22.5

require internal/pokecache v0.0.0

require internal/location v0.0.0

require internal/locationArea v0.0.0

require internal/response v0.0.0

require internal/pokemon v0.0.0

require internal/pokedex v0.0.0

replace internal/pokedex => ./internal/pokedex

replace internal/pokecache => ./internal/pokecache

replace internal/location => ./internal/location

replace internal/locationArea => ./internal/locationArea

replace internal/response => ./internal/response

replace internal/pokemon => ./internal/pokemon
