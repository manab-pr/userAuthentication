package main

import (
	"auth-project/config"
	"auth-project/graph"
	"auth-project/internal/middleware"
	"log"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := config.InitDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	r := gin.Default()

	// GraphQL setup
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &Resolver{DB: db}}))

	// Public endpoints
	r.POST("/query", gin.WrapH(srv))
	r.GET("/playground", gin.WrapH(playground.Handler("GraphQL playground", "/query")))

	// Protected endpoints
	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())
	protected.POST("/protected/query", gin.WrapH(srv))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on http://localhost:%s", port)
	log.Printf("GraphQL playground available at http://localhost:%s/playground", port)

	if err := r.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
