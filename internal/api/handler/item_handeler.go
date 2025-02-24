package handlers

import (
    "bidding-system/internal/services"
    "github.com/gin-gonic/gin"
    "net/http"
)

type ItemHandler struct {
    itemService *services.ItemService
}

func NewItemHandler(itemService *services.ItemService) *ItemHandler {
    return &ItemHandler{itemService: itemService}
}

// GetItems returns all available items for bidding
func (h *ItemHandler) GetItems(c *gin.Context) {
    items, err := h.itemService.GetItems()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"items": items})
}