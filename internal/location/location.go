package location

import (
	"encoding/json"
	"fmt"
	"internal/pokecache"
	"io"
	"net/http"
	"time"
)

var apiURL = "https://pokeapi.co/api/v2/"
var currentLocation = apiURL + "location/"
var previousLocation = ""
var nextLocation = currentLocation
var cache = pokecache.NewCache(time.Second * 20)

type LocationResponse struct {
	Count    int        `json:"count"`
	Next     string     `json:"next"`
	Previous string     `json:"previous"`
	Results  []Location `json:"results"`
}

func GetResponse(url string) ([]byte, error) {
	body, OK := cache.Get(url)

	if !OK {
		res, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			fmt.Println(res.Status)
			return nil, fmt.Errorf(res.Status)
		}

		body, err = io.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		cache.Add(url, body)
		//fmt.Println("Added to cache")
	}

	return body, nil
}

func unmarshalLocation(data []byte) (LocationResponse, error) {
	var resp LocationResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		fmt.Println(err)
		return LocationResponse{}, err
	}
	return resp, nil
}

type Location struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func GetLocations() ([]Location, error) {
	resp, err := GetResponse(currentLocation)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	data, err := unmarshalLocation(resp)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	currentLocation = data.Next
	previousLocation = data.Previous

	fmt.Println(currentLocation)

	return data.Results, nil
}

func GetPreviousLocations() ([]Location, error) {
	var err error
	prev := previousLocation
	if prev == "" {
		err = fmt.Errorf("No previous map found")
		fmt.Println(err)
		return nil, err
	}

	resp, err := GetResponse(prev)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	data, err := unmarshalLocation(resp)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	currentLocation = previousLocation
	previousLocation = data.Previous

	return data.Results, nil
}
