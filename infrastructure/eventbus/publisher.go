package eventbus

import (
	"github.com/redis/go-redis/v9"

	"redbank/data/event"
)

// Publisher is a Redis-backed implementation of event.Publisher.
type Publisher struct {
	client *redis.Client
}

func NewPublisher(client *redis.Client) *Publisher {
	return &Publisher{client: client}
}

func (p *Publisher) Publish(e event.Event) error {
	return nil
}
