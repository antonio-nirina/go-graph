package queries

import (
	"fmt"
	"io/ioutil"
	"errors"
	"log"
	"net/http"
	"encoding/json"
	//"reflect"

	// "github.com/antonio-nirina/go-graph/model"
	"github.com/antonio-nirina/go-graph/types"
	"github.com/graphql-go/graphql"
)

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data []User `json:"data"`
}

type ResponseSi struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data User `json:"data"`
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
var resp = User{}
var result = []User{}


// GetUserQuery returns the queries available against user type.
func GetUserQuery() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(types.UserType),
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			log.Printf("")
			fResp ,_ := fetchPost()
			// var users []types.User
			return *fResp, nil
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
				find , _ := fetchOnePost(idQuery)
				return *find,nil
			}

			return User{}, nil
		},
	}
}

func fetchPost() (*[]User, error) {
	var res = User{}
	var lastRes = []User{}
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

	result := Response{}
	err = json.Unmarshal(b, &result)
	if err != nil {
		return nil, errors.New("could not unmarshal data")
	}
	for _, todo := range result.Data {
		res.Addresse = todo.Addresse
		res.Avatar = todo.Avatar
		res.ID = todo.ID
		res.Phone = todo.Phone
		res.Email = todo.Email
		res.FirstName = todo.FirstName
		res.LastName = todo.LastName
		lastRes = append(lastRes,res)
	}

	return &lastRes, nil
}

func fetchOnePost(id string) (*User, error) {
	var res = User{}
	uri := fmt.Sprintf("%s%s","https://coursev1.herokuapp.com/api/user/",id)
	resp, err := http.Get(uri)

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

	result := ResponseSi{}
	err = json.Unmarshal(b, &result)

	if err != nil {
		return nil, errors.New("could not unmarshal data")
	}

	res.Addresse = result.Data.Addresse
	res.Avatar = result.Data.Avatar
	res.ID = result.Data.ID
	res.Phone = result.Data.Phone
	res.Email = result.Data.Email
	res.FirstName = result.Data.FirstName
	res.LastName = result.Data.LastName

	return &res, nil
}
