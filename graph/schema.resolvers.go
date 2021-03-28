package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/codehakase/buycoins-th/graph/generated"
)

func (r *queryResolver) CalculatePrice(ctx context.Context, typeArg string, margin float64, exchangeRate float64) (float64, error) {
	return r.PriceService.CalculatePrice(typeArg, margin, exchangeRate)
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
