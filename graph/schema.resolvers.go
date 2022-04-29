package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/json"

	"github.com/haulerkonj/gqlgen_todos/graph/generated"
	"github.com/haulerkonj/gqlgen_todos/graph/model"
)

func (r *mutationResolver) Fetch(ctx context.Context, input string) (*model.Record, error) {
	return r.records[input], nil
}

func (r *mutationResolver) Append(ctx context.Context, input model.NewRecord) (string, error) {
	r.records[input.Num] = &model.Record{
		Num:           input.Num,
		Names:         input.Names,
		SumOfAllForks: input.SumOfAllForks,
	}
	j, _ := json.Marshal(r.records)
	return string(j), nil
}

func (r *mutationResolver) Init(ctx context.Context) (*model.Record, error) {
	record := &model.Record{
		Num:           "0",
		Names:         "",
		SumOfAllForks: "",
		// User: &model.User{ID: input.UserID, Name: "user " + input.UserID},
	}
	// r.records = append(r.records, record)
	r.records = make(map[string]*model.Record)
	r.records["0"] = record
	return record, nil
}

func (r *queryResolver) Record(ctx context.Context, input string) (*model.Record, error) {
	return r.records[input], nil
}

func (r *queryResolver) Records(ctx context.Context) (string, error) {
	j, _ := json.Marshal(r.records)
	return string(j), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
