package graph

import (
	"auth-project/internal/models"
	"auth-project/internal/utils"
	"context"
	"errors"
	"strconv"

	"gorm.io/gorm"
)

type Resolver struct {
	DB *gorm.DB
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

func (r *mutationResolver) CreateUser(ctx context.Context, input CreateUserInput) (*models.User, error) {
	user := &models.User{
		Name:     input.Name,
		Email:    input.Email,
		Phone:    input.Phone,
		Password: input.Password,
		Role:     input.Role,
	}

	if err := user.HashPassword(); err != nil {
		return nil, err
	}

	if err := r.DB.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *mutationResolver) Login(ctx context.Context, input LoginInput) (*AuthResponse, error) {
	var user models.User
	if err := r.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		return nil, errors.New("invalid credentials")
	}

	if err := user.ComparePassword(input.Password); err != nil {
		return nil, errors.New("invalid credentials")
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return nil, err
	}

	return &AuthResponse{
		Token: token,
		User:  &user,
	}, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*models.User, error) {
	var users []*models.User
	if err := r.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *queryResolver) User(ctx context.Context, id string) (*models.User, error) {
	var user models.User
	userID, _ := strconv.ParseUint(id, 10, 64)
	if err := r.DB.First(&user, userID).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *queryResolver) Me(ctx context.Context) (*models.User, error) {
	userID := ctx.Value("user_id").(uint)
	var user models.User
	if err := r.DB.First(&user, userID).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
