package swa

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
)

const itemUrl = "https://www.steamwebapi.com/steam/api"

type ItemService struct {
	client *Client
}

func (c *Client) Items() (items *ItemService) {
	items = &ItemService{client: c}
	return
}

// GetAllItemsFromGame Discover a comprehensive database of all Steam Market items, encompassing prices, history, item names, item conditions, and a wealth of additional information.
// This dynamic item list provides endless possibilities for your programming needs.
// Allowed parameters: game, sort_by, search, item_group, item_type, item_name, wear, currency, no_cache, page, max, price_min, price_max
func (s *ItemService) GetAllItemsFromGame(gameName string, options ...interface{}) ([]*GameItem, error) {
	url := fmt.Sprintf("%s/%s", itemUrl, "items")

	// Create a map to hold the query parameters.
	queryParams := make(map[string]string)

	// Set the required parameter.
	queryParams["game"] = gameName

	// Define a list of allowed optional parameter names.
	allowedParams := []string{"sort_by", "search", "item_group", "item_type", "item_name", "wear", "currency", "no_cache", "page", "max", "price_min", "price_max"}

	// Process optional parameters.
	for i := 0; i < len(options); i += 2 {
		paramName, paramValue := options[i].(string), options[i+1]

		// Check if the provided parameter name is in the list of allowed parameters and if it's a string (for simplicity).
		if contains(allowedParams, paramName) {
			if str, ok := paramValue.(string); ok {
				queryParams[paramName] = str
			}
		}
	}

	body, err := s.client.Get(url, queryParams)
	if err != nil {
		return nil, err
	}

	var gameItems []*GameItem
	if err := json.Unmarshal(body, &gameItems); err != nil {
		return nil, err
	}

	return gameItems, nil
}

//steam/api/items

// SearchItems This API costs 1 credit per request. Search for items by game, type, or name.
// Allowed parameters: game, search, type
func (s *ItemService) SearchItems(gameName string, options ...string) ([]*SearchItem, error) {
	url := fmt.Sprintf("%s/%s/%s", itemUrl, "items", "find")

	// Create a map to hold the query parameters.
	queryParams := make(map[string]string)

	// Set the required parameter.
	queryParams["game"] = gameName

	// Define a list of allowed optional parameter names.
	allowedOptions := []string{"search", "type"}

	// Process optional parameters.
	for i := 0; i < len(options); i += 2 {
		paramName := options[i]
		paramValue := options[i+1]

		// Check if the provided parameter name is allowed.
		if contains(allowedOptions, paramName) {
			queryParams[paramName] = paramValue
		}
	}

	body, err := s.client.Get(url, queryParams)
	if err != nil {
		return nil, err
	}

	var searchItems []*SearchItem
	if err := json.Unmarshal(body, &searchItems); err != nil {
		return nil, err
	}

	return searchItems, nil
}

//steam/api/items/find

// GetItemDetail This API endpoint returns a single item by market_hash_name, slug or hashId. This API requires 1 credits per request. No Parameters.
// Allowed parameters: market_hash_name, currency
func (s *ItemService) GetItemDetail(marketHashName string, options ...string) (*ItemDetail, error) {
	url := fmt.Sprintf("%s/%s", itemUrl, "item")

	// Create a map to hold the query parameters.
	queryParams := make(map[string]string)

	// Set the required parameter.
	queryParams["market_hash_name"] = marketHashName

	// Process optional parameters.
	for i := 0; i < len(options); i += 2 {
		paramName := options[i]
		paramValue := options[i+1]

		// Check if the provided parameter name is "currency" and if it's a non-empty string.
		if paramName == "currency" {
			if paramValue != "" {
				queryParams["currency"] = paramValue
			}
		}
	}

	body, err := s.client.Get(url, queryParams)
	if err != nil {
		return nil, err
	}

	var itemDetail *ItemDetail
	if err := json.Unmarshal(body, &itemDetail); err != nil {
		return nil, err
	}

	return itemDetail, nil
}

//steam/api/item

// GetItemPriceHistory This API costs 1 credit per request. This API endpoint returns the price history for a single item by market_hash_name, slug or hashId.
// StartDate and EndDate are optional parameters. (format: YYYY-MM-DD)
func (s *ItemService) GetItemPriceHistory(marketHashName string, options ...string) ([]*ItemHistory, error) {
	url := fmt.Sprintf("%s/%s", itemUrl, "history")

	// Create a map to hold the query parameters.
	queryParams := make(map[string]string)

	// Set the required parameter.
	queryParams["market_hash_name"] = marketHashName

	// Define a list of allowed optional parameter names.
	allowedOptions := []string{"origin", "type", "source", "interval", "start_date", "end_date", "currency"}

	// Process optional parameters.
	for i := 0; i < len(options); i += 2 {
		paramName := options[i]
		paramValue := options[i+1]

		// Check if the provided parameter name is allowed.
		if contains(allowedOptions, paramName) {
			queryParams[paramName] = paramValue
		}
	}

	body, err := s.client.Get(url, queryParams)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	log.Info(string(body))

	var itemHistory []*ItemHistory
	if err := json.Unmarshal(body, &itemHistory); err != nil {
		log.Error("SA", err)
		return nil, err
	}

	return itemHistory, nil
}

//steam/api/item/history
