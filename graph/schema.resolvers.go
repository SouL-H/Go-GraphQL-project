package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"graphQLv2/graph/generated"
	"graphQLv2/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	todo := &model.Todo{
		ID:   fmt.Sprint(len(r.TodosList) + 1),
		Text: input.Text,
		Done: false,
		User: &model.User{
			ID:   input.UserID,
			Name: fmt.Sprintf("user%s", fmt.Sprint(len(r.TodosList)+1)),
		},
	}
	r.TodosList = append(r.TodosList, todo)
	return todo, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {

	return r.TodosList, nil
	//return []*model.Todo{
	// &model.Todo{
	// 	ID:   "1",
	// 	Text: "Todo 1",
	// 	Done:false,
	// 	User:&model.User{
	// 		ID:   "1",
	// 	},
	// },
	// },nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
