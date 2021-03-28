package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/codehakase/buycoins-th/graph"
	"github.com/codehakase/buycoins-th/graph/generated"
	"github.com/codehakase/buycoins-th/services/price"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	priceSrv := price.New()
	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{Resolvers: &graph.Resolver{priceSrv}},
		),
	)

	http.Handle("/", playground.Handler("GraphQL playground", "/qraphiql"))
	http.Handle("/qraphiql", srv)

	log.Printf("connect to http://localhost:%s/graphiql for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
