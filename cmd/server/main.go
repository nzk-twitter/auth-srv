package main

import (
    "github.com/gin-gonic/gin"
    "github.com/bidyutmondal/twitter-clone/auth-srv/internal/handlers"
    "github.com/bidyutmondal/twitter-clone/auth-srv/internal/repository"
    "github.com/bidyutmondal/twitter-clone/auth-srv/internal/service"
    "github.com/bidyutmondal/twitter-clone/auth-srv/pkg/database"
    "log"
)

func main() {
    db, err := database.InitDB("postgres://authuser:authpassword@localhost:5433/auth_db?sslmode=disable")
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }
    defer db.Close()

    userRepo := &repository.UserRepository{DB: db}
    authService := &service.AuthService{
        UserRepo:  userRepo,
        JWTSecret: "your-secret-key",
    }
    authHandler := &handlers.AuthHandler{AuthService: authService}

    r := gin.Default()

    r.POST("/register", authHandler.Register)
    r.POST("/login", authHandler.Login)

    if err := r.Run(":8081"); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}