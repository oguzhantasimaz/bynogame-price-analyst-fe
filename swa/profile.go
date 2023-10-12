package swa

import (
	"encoding/json"
	"fmt"
)

const profileUrl = "https://www.steamwebapi.com/explore/api"

type ProfileService struct {
	client *Client
}

func (c *Client) Profiles() (currencies *ProfileService) {
	currencies = &ProfileService{client: c}
	return
}

// GetRandomProfiles This API costs 2 credits per Profile.
// This API will return random 5 Profiles with Inventory information and Items, you can change the limit with the limit parameter.
func (s *ProfileService) GetRandomProfiles(limit int) (*RandomProfile, error) {
	url := fmt.Sprintf("%s/%s", profileUrl, "random")

	var queryParams map[string]string

	queryParams = map[string]string{
		"limit": fmt.Sprintf("%d", limit),
	}

	body, err := s.client.Get(url, queryParams)
	if err != nil {
		return nil, err
	}

	var randomProfile *RandomProfile
	if err := json.Unmarshal(body, &randomProfile); err != nil {
		return nil, err
	}

	return randomProfile, nil
}

// GetHighValueProfiles This API incurs a cost of 1 credit per profile.
// The API retrieves the top 5 profiles with the highest combined value across inventory, information, and items. You have the option to modify the limit using the 'limit' parameter.
func (s *ProfileService) GetHighValueProfiles(limit int) (*HighValueProfile, error) {
	url := fmt.Sprintf("%s/%s", profileUrl, "toplist")

	var queryParams map[string]string

	queryParams = map[string]string{
		"limit": fmt.Sprintf("%d", limit),
	}

	body, err := s.client.Get(url, queryParams)
	if err != nil {
		return nil, err
	}

	var highProfile *HighValueProfile
	if err := json.Unmarshal(body, &highProfile); err != nil {
		return nil, err
	}

	return highProfile, nil
}
