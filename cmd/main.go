package main

import (
	"github.com/abik1221/ONLINE_CHERETA/internal/api/handlers"
	"github.com/abik1221/ONLINE_CHERETA/internal/api/middleware"
	"github.com/abik1221/ONLINE_CHERETA/internal/config"
	"github.com/abik1221/ONLINE_CHERETA/internal/models"
	"github.com/abik1221/ONLINE_CHERETA/internal/repositories"
	"github.com/abik1221/ONLINE_CHERETA/internal/services"
	"log"

	"github.com/abik1221/ONLINE_CHERETA/internal/worker"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

    // Initialize Gin
	r := gin.Default()

	// Connect to the database
	dsn := "host=" + cfg.Database.Host + " user=" + cfg.Database.User + " password=" + cfg.Database.Password +
		" dbname=" + cfg.Database.DBName + " port=" + cfg.Database.Port + " sslmode=" + cfg.Database.SSLMode
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Auto-migrate models
	db.AutoMigrate(&models.User{}, &models.Item{}, &models.Bid{})

	// Initialize repositories
	itemRepo := repositories.NewItemRepository(db)
	bidRepo := repositories.NewBidRepository(db)

	// Initialize services
	itemService := services.NewItemService(itemRepo)
	bidService := services.NewBidService(bidRepo)

	// Initialize handlers
	itemHandler := handlers.NewItemHandler(itemService)
	bidHandler := handlers.NewBidHandler(bidService)

	// Protected routes
	authGroup := r.Group("/api")
	authGroup.Use(middleware.AuthMiddleware())
	{
		authGroup.GET("/profile", userHandler.GetUserProfile)
		authGroup.GET("/items", itemHandler.GetItems)
		authGroup.POST("/bids", bidHandler.PlaceBid)
		authGroup.GET("/bids", bidHandler.GetUserBids)
	}

	

	// Initialize WebSocket handler
	wsHandler := handlers.NewWebSocketHandler(bidService)

	// WebSocket route
	r.GET("/ws", func(c *gin.Context) {
		wsHandler.HandleWebSocket(c)
	})

     // HandleTeleBirrCallback handles the payment callback from TeleBirr
    func (h *TeleBirrHandler) HandleTeleBirrCallback(c *gin.Context) {
    type req struct {
        TransactionID string `json:"transaction_id" binding:"required"`
        Status        string `json:"status" binding:"required"` // e.g., "success", "failed"
    }

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Update the payment status in your database
    if req.Status == "success" {
        // Mark the payment as successful
    } else {
        // Mark the payment as failed
    }

    c.JSON(http.StatusOK, gin.H{"message": "callback received"})
}

    r.POST("/telebirr-callback", telebirrHandler.HandleTeleBirrCallback)

	go worker.StartWorker()

	// Start the server
	if err := r.Run(":" + cfg.Server.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
