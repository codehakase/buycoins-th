package price

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

var priceAPIURL = "https://api.coindesk.com/v1/bpi/currentprice.json"

// Service ...
type Service interface {
	CalculatePrice(typeArg string, margin, exchangeRate float64) (float64, error)
}

type srv struct {
	priceFunc func() (float64, error)
}

// New ...
func New() Service {
	return &srv{
		priceFunc: getCurrentPrice,
	}
}

// CalculatePrice ...
func (s *srv) CalculatePrice(typeArg string, margin, exchangeRate float64) (float64, error) {
	var newPrice float64

	if typeArg != "buy" && typeArg != "sell" {
		return 0, errors.New("Invalid type arg passed, allowed type: `buy` and `sell`")
	}

	percValue := margin / 100
	currentPrice, err := s.priceFunc()
	if err != nil {
		return 0, err
	}

	switch typeArg {
	case "sell":
		newPrice = currentPrice - (percValue * currentPrice)
	case "buy":
		newPrice = currentPrice + (percValue * currentPrice)
	}

	return (newPrice * exchangeRate), nil
}

// getCurrentPrice retrieves the current USD/BTC pair rate from the coindesk API
func getCurrentPrice() (float64, error) {
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest("GET", priceAPIURL, nil)
	if err != nil {
		return 0, err
	}

	req.Close = true
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return 0, err
	}

	// extract price data
	bpi := data["bpi"]
	bpiUSD := bpi.(map[string]interface{})["USD"]
	priceUSD := bpiUSD.(map[string]interface{})["rate_float"].(float64)

	return priceUSD, nil
}
