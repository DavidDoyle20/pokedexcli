package location

import (
	"encoding/json"
	"fmt"
	"internal/response"
)

var apiURL = "https://pokeapi.co/api/v2/"
var currentLocation = apiURL + "location/"
var previousLocation = ""
var nextLocation = currentLocation

type LocationResponse struct {
	Count    int        `json:"count"`
	Next     string     `json:"next"`
	Previous string     `json:"previous"`
	Results  []Location `json:"results"`
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
	Name  string `json:"name"`
	Url   string `json:"url"`
	Areas []Area `json:"areas"`
}
type Area struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func GetLocations() ([]Location, error) {
	resp, err := response.GetResponse(currentLocation)
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

	resp, err := response.GetResponse(prev)
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

func GetLocation(location string) (Location, error) {
	if location == "" {
		return Location{}, fmt.Errorf("No location provided!")
	}
	url := apiURL + "location/" + location

	resp, err := response.GetResponse(url)
	if err != nil {
		fmt.Println(err)
		return Location{}, err
	}
	data, err := unmarshalLocationArea(resp)
	if err != nil {
		fmt.Println(err)
		return Location{}, err
	}
	return data, nil
}
func unmarshalLocationArea(data []byte) (Location, error) {
	var resp Location
	if err := json.Unmarshal(data, &resp); err != nil {
		fmt.Println(err)
		return Location{}, err
	}
	return resp, nil
}
