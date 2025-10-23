package messenger

import "context"

type handleMessage func(adapterMessage) error

type adapterMessage struct {
	Queue      string
	Identifier string
	Body       string
}

// The adapter interface is used to communicate with the message broker.
type adapter interface {
	Dispatch(adapterMessage) error
	Subscribe(string, handleMessage, context.Context) error
}
