package queries

import (
	"fmt"
	"io/ioutil"
	"errors"
	"log"
	"net/http"
	"encoding/json"

	// "github.com/antonio-nirina/go-graph/model"
	"github.com/antonio-nirina/go-graph/types"
	"github.com/graphql-go/graphql"
)

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data []User `json:"data"`
}

type User struct {
	Addresse  string `json:"addresse"`
	Avatar    string `json:"avatar"`
	ID        string `json:"_id"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	
}

// var user = model.User{}
var resp = User{}
var result = []User{}
// GetUserQuery returns the queries available against user type.
func GetUserQuery() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(types.UserType),
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			log.Printf("")
			// var users []types.User
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
	}
}

func GetOneUserQuery() *graphql.Field {
	return &graphql.Field{
		Type:        types.UserType,
		Description: "Get single user",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {

			idQuery, isOK := params.Args["id"].(string)
			if isOK {
				// Search for el with id
				for _, todo := range TodoList {
					if todo.ID == idQuery {
						return todo, nil
					}
				}
			}

			return Todo{}, nil
		}
	}
}

func fetchPost() (*User, error) {
	resp, err := http.Get("https://coursev1.herokuapp.com/api/users")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("%s: %s", "could not fetch data", resp.Status)
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("could not read data")
	}
	
	result := User{}
	err = json.Unmarshal(b, &result)
	if err != nil {
		return nil, errors.New("could not unmarshal data")
	}
	fmt.Println(result)
	return &result, nil
}
