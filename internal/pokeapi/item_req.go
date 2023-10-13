package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListItems(pageURL *string) (ListItemResp, error) {
	endpoint := "/item"
	fullURL := baseURL + endpoint

	if pageURL != nil {
		fullURL = *pageURL
	}

	dat, ok := c.cache.Get(fullURL)
	if ok {
		fmt.Println("cache hit!")
		listItemResp := ListItemResp{}
		err := json.Unmarshal(dat, &listItemResp)
		if err != nil {
			return ListItemResp{}, err
		}

		return listItemResp, nil
	}
	fmt.Println("cache miss!")

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return ListItemResp{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return ListItemResp{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return ListItemResp{}, fmt.Errorf("bad status Code: %v", resp.StatusCode)
	}

	dat, err = io.ReadAll(resp.Body)
	if err != nil {
		return ListItemResp{}, err
	}

	locationAreasResp := ListItemResp{}
	err = json.Unmarshal(dat, &locationAreasResp)
	if err != nil {
		return ListItemResp{}, err
	}

	c.cache.Add(fullURL, dat)

	return locationAreasResp, nil

}

func (c *Client) GetItemList(LocationAreaName string) (LocationArea, error) {
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
