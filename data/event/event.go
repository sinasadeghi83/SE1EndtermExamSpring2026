package event

import "time"

// Event is the contract every domain event must satisfy to travel through
// the event-driven pipeline (publisher -> broker -> subscriber).
type Event interface {
	Name() string
	OccurredAt() time.Time
}

// Publisher is implemented by infrastructure adapters (e.g. Redis, Kafka)
// responsible for emitting domain events outside the process boundary.
type Publisher interface {
	Publish(event Event) error
}

// Subscriber is implemented by infrastructure adapters responsible for
// dispatching incoming domain events to registered handlers.
type Subscriber interface {
	Subscribe(eventName string, handler Handler) error
}

// Handler processes a single delivered Event.
type Handler func(Event) error
