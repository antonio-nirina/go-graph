package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/antonio-nirina/graph/queries"
	// "github.com/antonio-nirina/graph/mutations"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

func main() {
	schemaConfig := graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name:   "RootQuery",
			Fields: queries.GetRootFields(),
		}),
		/*Mutation: graphql.NewObject(graphql.ObjectConfig{
			Name:   "RootMutation",
			Fields: mutations.GetRootFields(),
		}),*/
	}
	schema, err := graphql.NewSchema(schemaConfig)

	if err != nil {
		log.Fatalf("Failed to create new schema, error: %v", err)
	}

	httpHandler := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	http.Handle("/", security.Handle(httpHandler))
	fmt.Println("ready: listening 4000")

	http.ListenAndServe(":4000", nil)
}
