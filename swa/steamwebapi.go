package swa

import (
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
)

const (
	BaseURL = "https://www.steamwebapi.com/"
)

type Client struct {
	APIKey  string
	BaseURL string
}

func NewClient(apiKey string) *Client {
	return &Client{
		APIKey:  apiKey,
		BaseURL: BaseURL,
	}
}

func (c *Client) GetAPIKey() string {
	return c.APIKey
}

func (c *Client) SetAPIKey(apiKey string) {
	c.APIKey = apiKey
}

func (c *Client) Get(url string, queryParameters map[string]string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	q := req.URL.Query()
	for k, v := range queryParameters {
		q.Add(k, v)
	}
	q.Add("key", c.APIKey)

	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	var body []byte
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return body, nil
}

// Not working
//	GetFloatValue This API incurs a cost of 1 Credit per request and provides float information for an item. You can provide parameters as follows: either inspectlink, or a combination of m, a, d, and s. At least one parameter is required.
//		func (s *SteamWebAPI) GetFloatValue(inspectLink, steamId, marketItemId, assetId, dId string) (string, error) {
//			return s.BaseURL + "GetFloatValue/v1/", nil
//	}
