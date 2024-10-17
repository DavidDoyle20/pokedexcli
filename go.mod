module github.com/DavidDoyle20/pokedexcli

go 1.22.5

require internal/pokecache v0.0.0

require internal/location v0.0.0

require internal/locationArea v0.0.0

replace internal/pokecache => ./internal/pokecache

replace internal/location => ./internal/location

replace internal/locationArea => ./internal/locationArea
