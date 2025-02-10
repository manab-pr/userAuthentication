package graph

import (
	"context"
	"strconv"
	"user-authentication/models"
	"user-authentication/services"
)

type Resolver struct {
	UserService *services.UserService
}

func (r *Resolver) GetUser(ctx context.Context, id string) (*models.User, error) {
	userID, _ := strconv.Atoi(id)
	return r.UserService.Getuser(uint(userID))
}

func (r *Resolver) ListUsers(ctx context.Context, name string, email string) (*models.User, error) {
	return r.UserService.CreateUser(name, email)
}

func (r *Resolver) Mutation() graph.Resolver {
	return &mutationResolver{r}
}
