package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (ListLocationAreas, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint

	if pageURL != nil {
		fullURL = *pageURL
	}

	dat, ok := c.cache.Get(fullURL)
	if ok {
		fmt.Println("cache hit!")
		locationAreasResp := ListLocationAreas{}
		err := json.Unmarshal(dat, &locationAreasResp)
		if err != nil {
			return ListLocationAreas{}, err
		}

		return locationAreasResp, nil
	}
	fmt.Println("cache miss!")

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return ListLocationAreas{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return ListLocationAreas{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return ListLocationAreas{}, fmt.Errorf("bad status Code: %v", resp.StatusCode)
	}

	dat, err = io.ReadAll(resp.Body)
	if err != nil {
		return ListLocationAreas{}, err
	}

	locationAreasResp := ListLocationAreas{}
	err = json.Unmarshal(dat, &locationAreasResp)
	if err != nil {
		return ListLocationAreas{}, err
	}

	c.cache.Add(fullURL, dat)

	return locationAreasResp, nil

}

func (c *Client) GetLocationArea(LocationAreaName string) (LocationArea, error) {
	endpoint := "/location-area/" + LocationAreaName
	fullURL := baseURL + endpoint

	dat, ok := c.cache.Get(fullURL)
	if ok {
		fmt.Println("cache hit!")
		locationArea := LocationArea{}
		err := json.Unmarshal(dat, &locationArea)
		if err != nil {
			return LocationArea{}, err
		}

		return locationArea, nil
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
		return LocationArea{}, fmt.Errorf("bad status Code: %v", resp.StatusCode)
	}

	dat, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationArea{}, err
	}

	locationArea := LocationArea{}
	err = json.Unmarshal(dat, &locationArea)
	if err != nil {
		return LocationArea{}, err
	}

	c.cache.Add(fullURL, dat)

	return locationArea, nil

}
