package handlers

import (
    "bidding-system/internal/services"
    "github.com/gin-gonic/gin"
    "net/http"
)

type BidHandler struct {
    bidService *services.BidService
}

func NewBidHandler(bidService *services.BidService) *BidHandler {
    return &BidHandler{bidService: bidService}
}

// PlaceBid places a bid on an item
func (h *BidHandler) PlaceBid(c *gin.Context) {
    userID := c.MustGet("userID").(uint)

    var req struct {
        ItemID    uint    `json:"item_id" binding:"required"`
        BidAmount float64 `json:"bid_amount" binding:"required"`
    }

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    bid, err := h.bidService.PlaceBid(userID, req.ItemID, req.BidAmount)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "bid placed", "bid": bid})
}

// GetUserBids returns all bids placed by the authenticated user
func (h *BidHandler) GetUserBids(c *gin.Context) {
    userID := c.MustGet("userID").(uint)

    bids, err := h.bidService.GetUserBids(userID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"bids": bids})
}