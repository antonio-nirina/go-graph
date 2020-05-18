package main

import (
	"fmt"
	"log"
	"net/http"

	// "github.com/antonio-nirina/go-graph/queries"
	// "github.com/antonio-nirina/go-graph/mutations"
	"github.com/antonio-nirina/go-graph/types"
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
var resp = User{}
var result = []User{} 
func main() {
	schemaConfig := graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name:   "RootQuery",
			// Fields: queries.GetRootFields(),
			Fields: graphql.Fields{
				"user":&graphql.Field{
					Type:        graphql.NewList(types.UserType),
					Description: "List of users",
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						resp.Addresse = ""
						resp.Avatar = ""
						resp.ID = "5eae6ac9d5aff51fbb501c4a"
						resp.Phone = "098734577"
						resp.Email = "zandry@gmail.com"
						resp.FirstName = "Jhon"
						resp.LastName = "Doe"
						result = append(result,resp) 
						return result, nil
					},
				},
			},
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
