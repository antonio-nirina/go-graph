package queries

import (
	"log"

	"github.com/antonio-nirina/go-graph/model"
	"github.com/antonio-nirina/go-graph/types"
	"github.com/graphql-go/graphql"
)

var user = model.User{}
// GetUserQuery returns the queries available against user type.
func GetUserQuery() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(types.UserType),
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			log.Printf("[query] user\n")
			var users []types.User

			return users, nil
		},
	}
}

func fetchPost() (string, error) {
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
	fmt.Println(string(b.data))
	return "eee", nil
}
