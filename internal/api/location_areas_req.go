package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreas, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint

	if pageURL != nil {
		fullURL = *pageURL
	}

	//check cache

	dat, ok := c.cache.Get(fullURL)
	if ok {
		//cache hit
		fmt.Println("cache hit!")
		locationAreasResp := LocationAreas{}
		err := json.Unmarshal(dat, &locationAreasResp)
		if err != nil {
			return LocationAreas{}, err
		}
		return locationAreasResp, nil
	}
	fmt.Println("cache miss!")

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreas{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreas{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationAreas{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	dat, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreas{}, err
	}

	locationAreasResp := LocationAreas{}
	err = json.Unmarshal(dat, &locationAreasResp)
	if err != nil {
		return LocationAreas{}, err
	}

	c.cache.Add(fullURL, dat)

	return locationAreasResp, nil
}

func (c *Client) GetLocationArea(locationAreaName string) (LocationArea, error) {
	endpoint := "/location-area/" + locationAreaName
	fullURL := baseURL + endpoint

	//check cache
	dat, ok := c.cache.Get(fullURL)
	if ok {
		//cache hit
		fmt.Println("cache hit!")
		locationAreasResp := LocationArea{}
		err := json.Unmarshal(dat, &locationAreasResp)
		if err != nil {
			return LocationArea{}, err
		}
		return locationAreasResp, nil
	}
	fmt.Println("cache miss!")

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationArea{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationArea{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	dat, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationArea{}, err
	}

	locationAreasResp := LocationArea{}
	err = json.Unmarshal(dat, &locationAreasResp)
	if err != nil {
		return LocationArea{}, err
	}

	c.cache.Add(fullURL, dat)

	return locationAreasResp, nil
}
