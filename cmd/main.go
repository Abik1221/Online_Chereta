package main

import (
    "ONLINE_CHARETA/internal/api/handlers"
    "ONLINE_CHARETA/internal/api/middleware"
    "ONLINE_CHARETA/internal/config"
    "ONLINE_CHARETA/internal/models"
    "ONLINE_CHARETA/internal/repositories"
    "ONLINE_CHARETA/internal/services"
    "github.com/gin-gonic/gin"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "log"
)

func main() {
    // Load configuration
    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatalf("Failed to load config: %v", err)
    }

    // Connect to the database
    dsn := "host=" + cfg.Database.Host + " user=" + cfg.Database.User + " password=" + cfg.Database.Password +
        " dbname=" + cfg.Database.DBName + " port=" + cfg.Database.Port + " sslmode=" + cfg.Database.SSLMode
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    // Auto-migrate models
    db.AutoMigrate(&models.User{})

    // Initialize repositories
    userRepo := repositories.NewUserRepository(db)

    // Initialize services
    userService := services.NewUserService(userRepo)

    // Initialize handlers
    userHandler := handlers.NewUserHandler(userService)

    // Initialize Gin
    r := gin.Default()

    // Public routes
    r.POST("/register", userHandler.RegisterUser)
    r.POST("/login", userHandler.LoginUser)

    // Protected routes
    authGroup := r.Group("/api")
    authGroup.Use(middleware.AuthMiddleware())
    {
        authGroup.GET("/profile", userHandler.GetUserProfile)
    }

    // Start the server
    if err := r.Run(":" + cfg.Server.Port); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}