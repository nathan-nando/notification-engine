package broker

import (
	"context"
	"encoding/json"
	"log"

	"github.com/redis/go-redis/v9"
)

type RedisBroker struct {
	client *redis.Client
}

func NewRedisBroker(redisURL string) (*RedisBroker, error) {
	opts, err := redis.ParseURL(redisURL)
	if err != nil {
		return nil, err
	}
	client := redis.NewClient(opts)
	
	// Check connection
	if err := client.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}
	
	return &RedisBroker{client: client}, nil
}

type ControlRequest struct {
	ChatID  int64    `json:"chat_id"`
	Command string   `json:"command"`
	Args    []string `json:"args"`
}

type ControlResponse struct {
	ChatID  int64  `json:"chat_id"`
	Message string `json:"message"`
}

func (r *RedisBroker) PublishControlRequest(ctx context.Context, req ControlRequest) error {
	payload, err := json.Marshal(req)
	if err != nil {
		return err
	}
	return r.client.Publish(ctx, "engine:control:requests", payload).Err()
}

func (r *RedisBroker) SubscribeControlResponses(ctx context.Context, handler func(ControlResponse)) {
	pubsub := r.client.Subscribe(ctx, "engine:control:responses")
	ch := pubsub.Channel()

	go func() {
		for msg := range ch {
			var resp ControlResponse
			if err := json.Unmarshal([]byte(msg.Payload), &resp); err != nil {
				log.Printf("Error unmarshalling control response: %v", err)
				continue
			}
			handler(resp)
		}
	}()
}
