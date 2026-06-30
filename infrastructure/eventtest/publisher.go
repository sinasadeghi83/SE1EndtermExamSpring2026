package eventtest

import (
	"sync"

	"redbank/data/event"
)

// Publisher is an in-memory event.Publisher implementation for unit tests.
// It records every published event so assertions can inspect them.
type Publisher struct {
	mu     sync.Mutex
	Events []event.Event
}

func New() *Publisher {
	return &Publisher{}
}

func (p *Publisher) Publish(e event.Event) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.Events = append(p.Events, e)
	return nil
}
