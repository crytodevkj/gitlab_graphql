/// [date] 2022-04-28

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/graphql-go/graphql"
)

type user struct {
	Num           string `json:"num"`
	Names         string `json:"names"`
	SumOfAllforks string `json:"sumOfAllForks"`
}

/*
   Create User object type with fields "num" and "res" by using GraphQLObjectTypeConfig:
       - Res: res of object type
       - Fields: a map of fields by using GraphQLFields
   Setup type of field use GraphQLFieldConfig
*/
var userType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"num": &graphql.Field{
				Type: graphql.String,
			},
			"names": &graphql.Field{
				Type: graphql.String,
			},
			"sumOfAllForks": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var _url string = ""
var _req string = ""
var _names string = ""
var _sumOfAllForks string = ""

/*
   Create Query object type with fields "user" has type [userType] by using GraphQLObjectTypeConfig:
       - Res: res of object type
       - Fields: a map of fields by using GraphQLFields
   Setup type of field use GraphQLFieldConfig to define:
       - Type: type of field
       - Args: arguments to query with current field
       - Resolve: function to query data using params from [Args] and return value with current type
*/
var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"user": &graphql.Field{
				Type: userType,
				Args: graphql.FieldConfigArgument{
					"num": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					idQuery, isOK := p.Args["num"].(string)
					if isOK {
						ret := user{}
						ret.Num = idQuery
						executeApi(idQuery)
						ret.Names = _names
						ret.SumOfAllforks = _sumOfAllForks
						return ret, nil
					}
					return nil, nil
				},
			},
		},
	})

var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query: queryType,
	},
)

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}
	return result
}

func executeApi(num string) {
	// call to repository layer
	_url = Url
	_req = strings.Replace(Req, "DISPLAY_NUM", num, -1)
	res := Api(_url, _req)
	// parse into array of node
	nodes := GetNodes(res)

	names := []string{}
	var sumOfAllForks float64 = 0
	// loop through array to get names | sum of all forks
	for i := 0; i < len(nodes); i++ {
		names = append(names, GetName(nodes[i]))
		sumOfAllForks += GetForksCount(nodes[i])
	}

	// display results
	_names = strings.Join(names, ", ")
	_sumOfAllForks = fmt.Sprint(sumOfAllForks)
}

func Service() {
	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		result := executeQuery(r.URL.Query().Get("query"), schema)
		json.NewEncoder(w).Encode(result)
	})

	fmt.Println("Now server is running on port 8080")
	fmt.Println("Test with Get      : curl -g 'http://localhost:8080/graphql?query={user(num:\"1\"){res}}'")
	http.ListenAndServe(":8080", nil)
}
