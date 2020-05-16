package types

import (

	// "github.com/antonio-nirina/graph/model"
	"github.com/graphql-go/graphql"
)

// User type definition.
type User struct {
	ID        int    `db:"id" json:"id"`
	Firstname string `db:"firstname" json:"firstname"`
	Lastname  string `db:"lastname" json:"lastname"`
	Email     string `db:"email" json:"email"`
	Password  string `db:"password" json:"password"`
	CreatedAt string `db:"created_at" json:"created_at"`
}

// UserType is the GraphQL schema for the user type.
var UserType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"_id":       &graphql.Field{Type: graphql.String},
		"firstName": &graphql.Field{Type: graphql.String},
		"lastName":  &graphql.Field{Type: graphql.String},
		"email":     &graphql.Field{Type: graphql.String},
		"phone":     &graphql.Field{Type: graphql.String},
		/*"roles": &graphql.Field{
			Type: graphql.NewList(RoleType),
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				var roles []Role

				// userID := params.Source.(User).ID
				// Implement logic to retrieve user associated roles from user id here.

				return roles, nil
			},
		},*/
	},
})
