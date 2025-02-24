package worker

import (
    "bidding-system/internal/utils"
    "context"
    "log"
    "github.com/hibiken/asynq"
)

func StartWorker() {
    redisAddr := "localhost:6379" // Replace with your Redis address
    srv := asynq.NewServer(
        asynq.RedisClientOpt{Addr: redisAddr},
        asynq.Config{Concurrency: 10},
    )

    mux := asynq.NewServeMux()
    mux.HandleFunc("send_notification", handleNotificationTask)

    if err := srv.Run(mux); err != nil {
        log.Fatalf("Failed to start worker: %v", err)
    }
}

func handleNotificationTask(ctx context.Context, t *asynq.Task) error {
    userID := t.Payload["user_id"].(float64)
    message := t.Payload["message"].(string)

    // Send notification via email or SMS
    log.Printf("Sending notification to user %d: %s", uint(userID), message)
    return nil
}