package swa

import (
	"encoding/json"
	"fmt"
)

const steamUrl = "https://www.steamwebapi.com/steam/api"

type SteamService struct {
	client *Client
}

func (c *Client) Steams() (steams *SteamService) {
	steams = &SteamService{client: c}
	return
}

// ConvertSteamId This API Endpoint converts a SteamID to SteamID2, SteamID3 and SteamID64 and returns it as JSON. Costs: 1 Credits
func (s *SteamService) ConvertSteamId(steamId string) (*SteamIds, error) {
	url := fmt.Sprintf("%s/%s", steamUrl, "steamid")

	var queryParams map[string]string

	queryParams = map[string]string{
		"id": steamId,
	}

	body, err := s.client.Get(url, queryParams)
	if err != nil {
		return nil, err
	}

	var sId *SteamIds
	if err := json.Unmarshal(body, &sId); err != nil {
		return nil, err
	}

	return sId, nil
}

//steam/api/steamid

// GetSteamProfile This endpoint allows you to fetch a Steam user's profile. The returned data is directly sourced from Steam without any caching. This API requires 2 credits per request.
func (s *SteamService) GetSteamProfile(steamId string, noCache ...int) (*SteamProfile, error) {
	url := fmt.Sprintf("%s/%s", steamUrl, "profile")

	var queryParams map[string]string

	queryParams = map[string]string{
		"id": steamId,
	}

	if len(noCache) > 0 {
		queryParams["no_cache"] = fmt.Sprintf("%d", noCache[0])
	}

	body, err := s.client.Get(url, queryParams)
	if err != nil {
		return nil, err
	}

	var sProfile *SteamProfile
	if err := json.Unmarshal(body, &sProfile); err != nil {
		return nil, err
	}

	return sProfile, nil
}

//steam/api/profile

// GetSteamFriends This endpoint allows you to fetch a Steam user's friends. The returned data is parsed by us, if you need directly sourced from Steam then send parse=0. Otherwise parsed=1 is default. This API requires 10 credits per request.
// Optional parameters: parse, noCache
func (s *SteamService) GetSteamFriends(steamId string, options ...interface{}) ([]*SteamFriends, error) {
	url := fmt.Sprintf("%s/%s", steamUrl, "friendlist")

	// Create a map to hold the query parameters.
	queryParams := make(map[string]string)

	// Set the required parameter.
	queryParams["id"] = steamId

	// Process optional parameters.
	for i := 0; i < len(options); i += 2 {
		paramName := options[i].(string)
		paramValue := options[i+1]

		// Check if the provided parameter name is "parse" or "noCache" and if it's an integer.
		if paramName == "parse" {
			if parse, ok := paramValue.(int); ok {
				queryParams["parse"] = fmt.Sprintf("%d", parse)
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

	var sFriends []*SteamFriends
	if err := json.Unmarshal(body, &sFriends); err != nil {
		return nil, err
	}

	return sFriends, nil
}

//steam/api/friendlist
