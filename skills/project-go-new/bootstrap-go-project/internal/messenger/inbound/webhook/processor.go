package webhook

import (
	"context"
)

// Processor handles webhooks for specific providers or types
type Processor interface {
	Supports(webhookType string) bool
	Process(ctx context.Context, msg *message) error
}
