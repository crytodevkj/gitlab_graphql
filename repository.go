package main

import (
	"context"

	"github.com/machinebox/graphql"
)

func Api(url string, req string) interface{} {
	graphqlClient := graphql.NewClient(url)
	graphqlRequest := graphql.NewRequest(req)
	var graphqlResponse interface{}

	if err := graphqlClient.Run(context.Background(), graphqlRequest, &graphqlResponse); err != nil {
		panic(err)
	}
	return graphqlResponse
}
