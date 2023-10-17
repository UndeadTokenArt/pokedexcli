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

func (c *Client) GetItemDetails(itemURL string) (ItemDetails, error) {
	endpoint := "/item/" + itemURL
	fullURL := baseURL + endpoint

	dat, ok := c.cache.Get(fullURL)
	if ok {
		fmt.Println("cache hit!")
		itemDetails := ItemDetails{}
		err := json.Unmarshal(dat, &itemDetails)
		if err != nil {
			return ItemDetails{}, err
		}

		return itemDetails, nil
	}
	fmt.Println("cache miss!")

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return ItemDetails{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return ItemDetails{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return ItemDetails{}, fmt.Errorf("bad status Code: %v", resp.StatusCode)
	}

	dat, err = io.ReadAll(resp.Body)
	if err != nil {
		return ItemDetails{}, err
	}

	locationArea := ItemDetails{}
	err = json.Unmarshal(dat, &locationArea)
	if err != nil {
		return ItemDetails{}, err
	}

	c.cache.Add(fullURL, dat)

	return ItemDetails{}, nil

}
