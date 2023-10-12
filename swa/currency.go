package swa

import (
	"encoding/json"
	"fmt"
)

const currencyUrl = "https://www.steamwebapi.com/currency/api"

type CurrencyService struct {
	client *Client
}

func (c *Client) Currencies() (currencies *CurrencyService) {
	currencies = &CurrencyService{client: c}
	return
}

// GetCurrencyList This API costs 5 credits per request.
// This API endpoint provides a list of all available currencies that can be used with the "base" parameter for currency conversion.
// By default, the base currency is set to USD (United States Dollar).
func (s *CurrencyService) GetCurrencyList(base string) (*CurrencyList, error) {
	url := fmt.Sprintf("%s/%s", currencyUrl, "list")

	var queryParams map[string]string

	if base != "" {
		queryParams = map[string]string{
			"base": base,
		}
	}

	body, err := s.client.Get(url, queryParams)
	if err != nil {
		return nil, err
	}

	var currencyList *CurrencyList
	if err := json.Unmarshal(body, &currencyList); err != nil {
		return nil, err
	}

	return currencyList, nil
}

//currency/api/list

// GetExchangeRate This API costs 1 credit per request.
// It returns an exchange rate for the specified currency. The default exchange rate is based on the USD exchange rate, which you can change using the "base" parameter.
// You can specify the currency you want to exchange to using the "change" parameter. For example, "change=EUR" and "base=USD" means you want to exchange from EUR to USD.
func (s *CurrencyService) GetExchangeRate(base, change string) (*ExchangeRate, error) {
	url := fmt.Sprintf("%s/%s", currencyUrl, "exchange")

	var queryParams map[string]string

	queryParams = map[string]string{
		"base":   base,
		"change": change,
	}

	body, err := s.client.Get(url, queryParams)
	if err != nil {
		return nil, err
	}

	var exchangeRate *ExchangeRate
	if err := json.Unmarshal(body, &exchangeRate); err != nil {
		return nil, err
	}

	return exchangeRate, nil
}

//currency/api/exchange
