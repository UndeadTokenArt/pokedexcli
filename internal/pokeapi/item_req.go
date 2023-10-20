package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListItems(pageURL *string) (ListItemsResp, error) {
	endpoint := "/item"
	fullURL := baseURL + endpoint

	if pageURL != nil {
		fullURL = *pageURL
	}

	dat, ok := c.cache.Get(fullURL)
	if ok {
		fmt.Println("cache hit!")
		listItemResp := ListItemsResp{}
		err := json.Unmarshal(dat, &listItemResp)
		if err != nil {
			return ListItemsResp{}, err
		}

		return listItemResp, nil
	}
	fmt.Println("cache miss!")

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return ListItemsResp{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return ListItemsResp{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return ListItemsResp{}, fmt.Errorf("bad status Code: %v", resp.StatusCode)
	}

	dat, err = io.ReadAll(resp.Body)
	if err != nil {
		return ListItemsResp{}, err
	}

	listItemResp := ListItemsResp{}
	err = json.Unmarshal(dat, &listItemResp)
	if err != nil {
		return ListItemsResp{}, err
	}

	c.cache.Add(fullURL, dat)
	return listItemResp, nil

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
		fmt.Println("newRequest")
		return ItemDetails{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		fmt.Println("httpClientDo")
		return ItemDetails{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		fmt.Println("StatusCode")
		return ItemDetails{}, fmt.Errorf("bad status Code: %v", resp.StatusCode)
	}

	dat, err = io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ReadAll")
		return ItemDetails{}, err
	}

	itemDetails := ItemDetails{}
	err = json.Unmarshal(dat, &itemDetails)
	if err != nil {
		fmt.Println("Unmarshal")
		return ItemDetails{}, err
	}

	c.cache.Add(fullURL, dat)
	fmt.Println("clean")
	return itemDetails, nil

}
