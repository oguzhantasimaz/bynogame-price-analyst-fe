package swa

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
)

const inventoryUrl = "https://www.steamwebapi.com/steam/api/inventory"

type InventoryService struct {
	client *Client
}

func (c *Client) Inventories() (inventories *InventoryService) {
	inventories = &InventoryService{client: c}
	return
}

// GetSteamInventory This endpoint utilizes a pool of proxies to fetch the Steam inventory data, allowing us to bypass rate-limiting restrictions. The data returned is an exact representation of Steam's response, ensuring accuracy and real-time information.
// However, if you prefer a more user-friendly format, you can enable the parse parameter. This will provide a parsed version of the data, including item prices and comprehensive details. It's important to note that each request to this API consumes 5 credits.
func (s *InventoryService) GetSteamInventory(steamId, gameName string, options ...interface{}) ([]*SteamInventory, error) {
	url := fmt.Sprintf("%s", inventoryUrl)

	// Create a map to hold the query parameters.
	queryParams := make(map[string]string)

	// Set the required parameters.
	queryParams["steam_id"] = steamId
	queryParams["game"] = gameName

	// Process optional parameters.
	for i := 0; i < len(options); i += 2 {
		paramName := options[i].(string)
		paramValue := options[i+1]

		// Check if the provided parameter name is one of the optional parameters.
		if paramName == "language" {
			if str, ok := paramValue.(string); ok {
				queryParams["language"] = str
			}
		} else if paramName == "parse" {
			if parse, ok := paramValue.(int); ok {
				queryParams["parse"] = fmt.Sprintf("%d", parse)
			}
		} else if paramName == "group" {
			if group, ok := paramValue.(int); ok {
				queryParams["group"] = fmt.Sprintf("%d", group)
			}
		} else if paramName == "noCache" {
			if noCache, ok := paramValue.(int); ok {
				queryParams["no_cache"] = fmt.Sprintf("%d", noCache)
			}
		}
	}

	body, err := s.client.Get(url, queryParams)
	if err != nil {
		return nil, err
	}

	log.Println(string(body))

	var steamInventories []*SteamInventory
	if err := json.Unmarshal(body, &steamInventories); err != nil {
		return nil, err
	}

	return steamInventories, nil
}

//steam/api/inventory

// GetSteamInventoryAtTime Get the inventory of a steam profile for a specific day. The inventory is saved and can be retrieved from there. But before you need to call the inventory endpoint.
func (s *InventoryService) GetSteamInventoryAtTime(steamId, date string) ([]*SteamInventory, error) {
	url := fmt.Sprintf("%s/%s/%s", inventoryUrl, "tracked", "day")

	var queryParams map[string]string

	queryParams = map[string]string{
		"steam_id": steamId,
		"date":     date,
	}

	body, err := s.client.Get(url, queryParams)
	if err != nil {
		return nil, err
	}

	var steamInventories []*SteamInventory
	if err := json.Unmarshal(body, &steamInventories); err != nil {
		return nil, err
	}

	return steamInventories, nil
}

//steam/api/inventory/tracked/day
