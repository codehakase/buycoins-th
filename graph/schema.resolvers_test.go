package graph

import (
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/codehakase/buycoins-th/graph/generated"
	"github.com/codehakase/buycoins-th/services/price/mocks"
	"github.com/stretchr/testify/mock"
)

func TestQueryPriceCalculation(t *testing.T) {
	t.Run("should return correct price calculation", func(t *testing.T) {
		priceSrv := new(mocks.Service)
		resolvers := Resolver{PriceService: priceSrv}
		c := client.New(handler.NewDefaultServer(
			generated.NewExecutableSchema(
				generated.Config{Resolvers: &resolvers},
			),
		))
		priceCal := 27851852.122
		priceSrv.On("CalculatePrice", mock.Anything, mock.AnythingOfType("float64"), mock.AnythingOfType("float64")).
			Return(priceCal, nil)
		var resp map[string]interface{}
		q := `
		query CalculatePrice {
			calculatePrice(type: "buy", margin:0.2, exchangeRate:476)
		  } 
	`
		c.MustPost(q, &resp)
		priceSrv.AssertExpectations(t)
		if resp["calculatePrice"] != priceCal {
			t.Errorf("expected priceCal be %v got %v", priceCal, resp["calculatePrice"])
		}
	})
}
