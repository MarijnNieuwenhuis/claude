package webhook

import (
	"context"
	"encoding/json"

	"gitlab.com/btcdirect-api/go-modules/messenger"
	"go.uber.org/zap"
)

type handler struct {
	processors []Processor
	logger     *zap.SugaredLogger
}

// NewHandler creates a new webhook message handler
func NewHandler(
	processors []Processor,
	logger *zap.SugaredLogger,
) messenger.MessageHandler {
	return &handler{
		processors: processors,
		logger:     logger,
	}
}

// Message implements messenger.MessageHandler
func (h *handler) Message() messenger.Message {
	return &message{}
}

// Handle implements messenger.MessageHandler
func (h *handler) Handle(m messenger.Message) error {
	msg := m.(*message)
	ctx := context.Background()

	// Dispatch to appropriate processor
	for _, processor := range h.processors {
		if processor.Supports(msg.Payload.Type) {
			return processor.Process(ctx, msg)
		}
	}

	// No processor found for this webhook type
	h.logger.Debugw("No processor found for webhook type", "type", msg.Payload.Type)
	return nil
}

// WebhookPayload represents a generic webhook payload structure
type WebhookPayload struct {
	Type string                 `json:"type"`
	Data map[string]interface{} `json:"data"`
}

type message struct {
	Headers    map[string]string `json:"-"`
	Payload    WebhookPayload    `json:"payload"`
	RawPayload string            `json:"-"` // For signature validation
}

func (m *message) Queue() string {
	return "bootstrap-go-service.webhook"
}

func (m *message) Identifier() string {
	return "webhook"
}

func (m *message) UnmarshalJSON(data []byte) error {
	var body struct {
		Headers map[string]string `json:"headers"`
		Payload string            `json:"payload"`
	}
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}

	m.Headers = body.Headers
	m.RawPayload = body.Payload
	return json.Unmarshal([]byte(body.Payload), &m.Payload)
}
