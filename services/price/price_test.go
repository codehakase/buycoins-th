package price

import (
	"fmt"
	"testing"
	//"github.com/codehakase/buycoins-th/services/price/mocks"
)

func TestInvalidTypeArg(t *testing.T) {
	s := New()
	_, err := s.CalculatePrice("random", 0.03, 390)
	if err == nil {
		t.Error("expected invalid typeArg to be caught")
	}
}

func TestCalculatePrice(t *testing.T) {
	s := New()
	testcases := []struct {
		cusPrice             float64
		_type                string
		margin, exchangeRate float64
		expected             float64
	}{
		{55888.1367, "sell", 0.33, 500, 27851852.924445},
		{56166.1467, "sell", 0.33, 500, 2.7990399207945e+07},
		{56166.1467, "buy", 0.33, 500, 2.8175747492055e+07},
		{56180.9083, "sell", 0.2, 476, 2.66886281260984e+07},
		{56180.9083, "buy", 0.2, 476, 2.6795596575501602e+07},
	}

	for _, tt := range testcases {
		t.Run(fmt.Sprintf("%s price for %v", tt._type, tt.cusPrice), func(t *testing.T) {
			s.(*srv).priceFunc = customPriceFunc(tt.cusPrice)
			price, err := s.CalculatePrice(tt._type, tt.margin, tt.exchangeRate)
			if err != nil {
				t.Errorf("expected nil errors, got: %v", err)
			}
			if price != tt.expected {
				t.Errorf("expected %v got %v", tt.expected, price)
			}
		})
	}
}

func customPriceFunc(basePrice float64) func() (float64, error) {
	return func() (float64, error) { return basePrice, nil }
}
