package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/antonio-nirina/go-graph/queries"
	// "github.com/antonio-nirina/go-graph/mutations"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

type User struct {
	Addresse  string `json:"addresse"`
	Avatar    string `json:"avatar"`
	ID        string `json:"_id"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	
}
// var resp = User{}
// var result = []User{} 
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

	// http.Handle("/", security.Handle(httpHandler))
	http.Handle("/",httpHandler)
	fmt.Println("ready: listening 4000")
	http.ListenAndServe(":4000", nil)
}
