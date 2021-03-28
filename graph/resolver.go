package graph

import "github.com/codehakase/buycoins-th/services/price"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver ...
type Resolver struct {
	PriceService price.Service
}
