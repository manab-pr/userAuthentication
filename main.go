package main

import (
	"user-authentication/config"
	"user-authentication/graph"
	"user-authentication/handlers"
	"user-authentication/repositories"
	"user-authentication/services"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDB()

	userRepo := repositories.NewUserRepository(config.DB)
	userService := services.NewUserService(userRepo)
	resolver := &graph.Resolver{UserService: userService}

	r := gin.Default()

	r.POST("/graphql", handlers.GraphQLHandler(resolver))
	r.GET("/", handlers.PlaygroundHandler())

	r.Run(":8080")
}
