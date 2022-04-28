/// [date] 2022-04-28
/// [ref] https://www.thepolyglotdeveloper.com/2020/02/interacting-with-a-graphql-api-with-golang/

package main

import (
	"context"

	"github.com/machinebox/graphql"
)

/// @param url: URL to graphql server
/// @param req: graphql query formatted into SDL
/// @return: graphql response
func Api(url string, req string) interface{} {
	graphqlClient := graphql.NewClient(url)
	graphqlRequest := graphql.NewRequest(req)
	var graphqlResponse interface{}

	// execute api
	if err := graphqlClient.Run(context.Background(), graphqlRequest, &graphqlResponse); err != nil {
		panic(err)
	}
	return graphqlResponse
}
