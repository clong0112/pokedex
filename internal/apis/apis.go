package apis

import (
	"encoding/json"
	"io"
	"net/http"
	"fmt"
	"github.com/clong0112/pokedex/internal/pokecache"
)


type LocationAreas struct {
		Next *string 
		Previous *string			
		Results []NameUrl
}

type NameUrl struct {
		Name string
		URL string
}

func GetLocationArea(url string, c *Config) (LocationAreas, error) {
	body, ok := c.Cache.Get(url)
	if ok {
		fmt.Println("Cache hit!")
		var data LocationAreas
		err := json.Unmarshal(body, &data)
		if err != nil {
			return LocationAreas{}, err
		}
		return data, nil
	}

	fmt.Println("Cache miss!")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreas{}, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return LocationAreas{}, err
	}
	defer resp.Body.Close()

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreas{}, err
	}

	c.Cache.Add(url, body)

	var data LocationAreas
	err = json.Unmarshal(body, &data)
	if err != nil {
		return LocationAreas{}, err
	}

	return data, nil

}

type CliCommand struct {
	Name string
	Description string
	Callback func(*Config) error
}

type Config struct {
	Next string
	Previous string
	Cache *pokecache.Cache
}