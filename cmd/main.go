package main

import (
	"fitness/config"
	"fitness/internal/database"
	"fitness/internal/handler"
	"fitness/internal/repository"
	"fitness/internal/service"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := database.ConnectPostgres()
	if err != nil {
		log.Fatal("Could not connect to the database:", err)
	}

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	authHandler := handler.AuthHandler{UserService: userService}
	userHandler := handler.UserHandler{UserService: userService}

	router := gin.Default()

	router.POST("/api/login", authHandler.Login)

	router.POST("/api/users", userHandler.CreateUser)
	router.GET("/api/users/:id", userHandler.GetUserByID)
	router.PUT("/api/users/:id", userHandler.UpdateUser)
	router.DELETE("/api/users/:id", userHandler.DeleteUser)

	port := config.GetEnv("PORT", "8080")
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Could not run server:", err)
	}
}
