package action

import (
	"fmt"

	"gitlab.com/btcdirect-api/go-modules/messenger"
	"go.uber.org/zap"
)

// Event represents a generic event to be published
type Event struct {
	Type string                 `json:"type"`
	Data map[string]interface{} `json:"data"`
}

// messageDispatcher defines the interface for dispatching messages
type messageDispatcher interface {
	Dispatch(msg messenger.Message) error
}

// Publisher publishes event messages
type Publisher struct {
	messenger messageDispatcher
	logger    *zap.SugaredLogger
}

// NewPublisher creates a new event publisher
func NewPublisher(messenger messenger.Messenger, logger *zap.SugaredLogger) *Publisher {
	return &Publisher{
		messenger: messenger,
		logger:    logger,
	}
}

// PublishEvent publishes an event
func (p *Publisher) PublishEvent(event Event, queue string) error {
	msg := &eventMessage{
		Type:  event.Type,
		Data:  event.Data,
		queue: queue,
	}

	p.logger.Infow("Publishing event message",
		"type", msg.Type,
		"queue", queue,
	)

	if err := p.messenger.Dispatch(msg); err != nil {
		return fmt.Errorf("failed to dispatch event message: %w", err)
	}

	return nil
}

// eventMessage represents a generic event notification
type eventMessage struct {
	Type  string                 `json:"type"`
	Data  map[string]interface{} `json:"data"`
	queue string
}

// Queue implements messenger.Message
func (m *eventMessage) Queue() string {
	return m.queue
}

// Identifier implements messenger.Message
func (m *eventMessage) Identifier() string {
	return "event"
}
