package model

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import "time"

// User represents the user model
type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// AuthResponse represents the authentication response
type AuthResponse struct {
	Token string `json:"token"`
	User  *User  `json:"user"`
}

// LoginInput represents login input data
type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// RegisterInput represents registration input data
type RegisterInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
