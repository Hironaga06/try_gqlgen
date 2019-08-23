package try_gqlgen

import (
	"context"
	"fmt"
	"math/rand"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct {
	users []*User
	todos []*Todo
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) User() UserResolver {
	return &userResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateUser(ctx context.Context, input NewUser) (*User, error) {
	user := &User{
		ID:   fmt.Sprintf("T%d", rand.Int()),
		Name: input.Name,
		Age:  input.Age,
	}
	r.users = append(r.users, user)
	return user, nil
}
func (r *mutationResolver) CreateTodo(ctx context.Context, input NewTodo) (*Todo, error) {
	todo := &Todo{
		ID:     fmt.Sprintf("T%d", rand.Int()),
		UserID: input.UserID,
		Text:   input.Text,
	}
	r.todos = append(r.todos, todo)
	return todo, nil
}

func (r *mutationResolver) UpdateTodo(ctx context.Context, input UpdateTodo) (*Todo, error) {
	var todo *Todo
	for _, v := range r.todos {
		if v.ID == input.ID {
			v.Done = input.Done
			todo = v
			break
		}
	}
	return todo, nil
}

func (r *mutationResolver) DeleteTodo(ctx context.Context, input DeleteTodo) (string, error) {
	todos := []*Todo{}
	for _, v := range r.todos {
		if v.ID != input.ID {
			todos = append(todos, v)
		}
	}
	r.todos = todos
	return "OK", nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Users(ctx context.Context) ([]*User, error) {
	return r.users, nil
}
func (r *queryResolver) Todos(ctx context.Context) ([]*Todo, error) {
	return r.todos, nil
}

type userResolver struct{ *Resolver }

func (r *userResolver) Todos(ctx context.Context, obj *User, isAll bool) ([]*Todo, error) {
	todos := []*Todo{}
	for _, v := range r.todos {
		if v.UserID != obj.ID {
			continue
		}
		todos = append(todos, v)
	}
	if isAll {
		return todos, nil
	}

	notExecutedTodos := []*Todo{}
	for _, v := range todos {
		if v.Done {
			continue
		}
		notExecutedTodos = append(notExecutedTodos, v)
	}
	return notExecutedTodos, nil
}
