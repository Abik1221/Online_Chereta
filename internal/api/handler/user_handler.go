package handlers

import (
    "ONLINE_CHARETA/internal/services"
    "ONLINE_CHARETA/internal/utils"
    "github.com/gin-gonic/gin"
    "net/http"
)

type UserHandler struct {
    userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
    return &UserHandler{userService: userService}
}

// RegisterUser handles user registration


// GetUserProfile returns the authenticated user's profile
func (h *UserHandler) GetUserProfile(c *gin.Context) {
    userID := c.MustGet("userID").(uint)

    // Fetch the user from the database
    user, err := h.userService.GetUserByID(userID)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"user": user})
}

func (h *UserHandler) RegisterUser(c *gin.Context) {
    var req struct {
        Username string `json:"username" binding:"required"`
        Email    string `json:"email" binding:"required,email"`
        Password string `json:"password" binding:"required,min=6"`
    }

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user, err := h.userService.RegisterUser(req.Username, req.Email, req.Password)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "user registered", "user": user})
}

// LoginUser handles user login
func (h *UserHandler) LoginUser(c *gin.Context) {
    var req struct {
        Email    string `json:"email" binding:"required,email"`
        Password string `json:"password" binding:"required,min=6"`
    }

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user, err := h.userService.LoginUser(req.Email, req.Password)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
        return
    }

    // Generate JWT token
    token, err := utils.GenerateJWT(user.ID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "login successful", "token": token})
}