package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) LocationDeepInfo(Location string) (LocationInfo, error) {

	url := baseURL + "/location-area/" + Location

	// cache check
	if val, ok := c.cache.Get(url); ok {
		locationInfo := LocationInfo{}
		err := json.Unmarshal(val, &locationInfo)
		if err != nil {
			return LocationInfo{}, err
		}
		return locationInfo, nil
	}

	// request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationInfo{}, err
	}

	// pull response data from request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationInfo{}, err
	}
	defer resp.Body.Close()

	// read data
	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationInfo{}, err
	}

	// unmarshal data from JSON
	locationInfo := LocationInfo{}
	err = json.Unmarshal(dat, &locationInfo)
	if err != nil {
		return LocationInfo{}, err
	}

	//add data to cache, return
	c.cache.Add(url, dat)
	return locationInfo, nil
}
