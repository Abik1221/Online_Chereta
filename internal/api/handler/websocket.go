package handlers

import (
    "bidding-system/internal/services"
    "github.com/gorilla/websocket"
    "log"
    "net/http"
)

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool {
        return true // Allow all connections (customize for production)
    },
}

type WebSocketHandler struct {
    bidService *services.BidService
}

func NewWebSocketHandler(bidService *services.BidService) *WebSocketHandler {
    return &WebSocketHandler{bidService: bidService}
}

// HandleWebSocket handles WebSocket connections for live bidding
func (h *WebSocketHandler) HandleWebSocket(c *gin.Context) {
    conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
    if err != nil {
        log.Println("WebSocket upgrade failed:", err)
        return
    }
    defer conn.Close()

    for {
        // Read message from the client
        var msg struct {
            ItemID    uint    `json:"item_id"`
            BidAmount float64 `json:"bid_amount"`
        }
        if err := conn.ReadJSON(&msg); err != nil {
            log.Println("WebSocket read error:", err)
            break
        }

        // Place the bid
        userID := c.MustGet("userID").(uint)
        bid, err := h.bidService.PlaceBid(userID, msg.ItemID, msg.BidAmount)
        if err != nil {
            conn.WriteJSON(gin.H{"error": err.Error()})
            continue
        }

        // Broadcast the new bid to all connected clients
        broadcastMessage := gin.H{
            "item_id":    msg.ItemID,
            "bid_amount": bid.BidAmount,
            "user_id":    bid.UserID,
        }
        conn.WriteJSON(broadcastMessage)
    }
}