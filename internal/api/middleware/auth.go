package middleware

import (
    "ONLINE_CHARETA/internal/utils"
    "github.com/gin-gonic/gin"
    "net/http"
)

// AuthMiddleware verifies the JWT token and sets the user ID in the context
func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenString := c.GetHeader("Authorization")
        if tokenString == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "authorization header required"})
            c.Abort()
            return
        }

        claims, err := utils.VerifyJWT(tokenString)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
            c.Abort()
            return
        }

        // Set the user ID in the context
        userID := uint(claims["user_id"].(float64))
        c.Set("userID", userID)

        c.Next()
    }
}