package services

import (
    "bidding-system/internal/models"
    "bidding-system/internal/utils"
    "github.com/hibiken/asynq"
)

type NotificationService struct {
    client *asynq.Client
}

func NewNotificationService(client *asynq.Client) *NotificationService {
    return &NotificationService{client: client}
}

// SendBidNotification sends a notification about a bid update
func (s *NotificationService) SendBidNotification(userID uint, message string) error {
    task := asynq.NewTask("send_notification", map[string]interface{}{
        "user_id": userID,
        "message": message,
    })
    _, err := s.client.Enqueue(task)
    return err
}