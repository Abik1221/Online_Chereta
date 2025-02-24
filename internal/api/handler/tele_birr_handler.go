package handlers

import (
    "bidding-system/internal/services"
    "github.com/gin-gonic/gin"
    "net/http"
)

type TeleBirrHandler struct {
    telebirrService *services.TeleBirrService
}

func NewTeleBirrHandler(telebirrService *services.TeleBirrService) *TeleBirrHandler {
    return &TeleBirrHandler{telebirrService: telebirrService}
}

// InitiateTeleBirrPayment initiates a payment via TeleBirr
func (h *TeleBirrHandler) InitiateTeleBirrPayment(c *gin.Context) {
    var req struct {
        Amount      float64 `json:"amount" binding:"required"`
        PhoneNumber string  `json:"phone_number" binding:"required"`
    }

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Call TeleBirr service to initiate payment
    transactionID, err := h.telebirrService.InitiatePayment(req.Amount, req.PhoneNumber, "https://your-callback-url.com/callback")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"transaction_id": transactionID})
}