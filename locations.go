package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var apiURL = "https://pokeapi.co/api/v2/"
var currentLocation = apiURL + "location/"
var previousLocation = ""
var nextLocation = currentLocation

type Response struct {
	Count    int        `json:"count"`
	Next     string     `json:"next"`
	Previous string     `json:"previous"`
	Results  []Location `json:"results"`
}

func getResponse(url string) (Response, error) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return Response{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println(res.Status)
		return Response{}, fmt.Errorf(res.Status)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return Response{}, err
	}

	var resp Response
	if err := json.Unmarshal(body, &resp); err != nil {
		fmt.Println(err)
		return Response{}, err
	}
	return resp, nil
}

type Location struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func getLocations(r Response) ([]Location, error) {
	resp, err := getResponse(currentLocation)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return resp.Results, nil

}

func getPreviousLocations() ([]Location, error) {
	resp, err := getResponse(currentLocation)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	prev := resp.Previous
	if prev == "" {
		err = fmt.Errorf("No previous map found")
		fmt.Println(err)
		return nil, err
	}

	resp, err = getResponse(prev)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return nil, nil
}
