package messenger

import (
	"encoding/json"
	"errors"
	"time"

	"gitlab.com/btcdirect-api/go-modules/app"
	"go.uber.org/zap"
)

type Config struct {
	Log            *zap.SugaredLogger
	Shutdown       *app.GracefulShutdown
	Environment    string
	RestartTimeout time.Duration
	PubsubConfig
}

type Messenger interface {
	Dispatch(Message) error
	Subscribe(...MessageHandler) error
}

type MessageDispatcher interface {
	Dispatch(Message) error
}

type Message interface {
	Identifier() string
	Queue() string
}

// Make sure to return the message by reference in order to be able to unmarshal it.
type MessageHandler interface {
	Message() Message
	Handle(Message) error
}

type messenger struct {
	Config
	adapter adapter
}

var ErrDifferentQueues = errors.New("all handlers must subscribe to the same queue")

// Creates a messenger instance using the Pub/Sub adapter.
// This also opens a connection to the message broker.
func New(c Config) Messenger {
	c.Log.Info("Starting messenger")
	c.PubsubConfig.DeadLetterTopic = c.Environment + "." + c.PubsubConfig.DeadLetterTopic
	a, err := newPubsubAdapter(c.PubsubConfig, c.Log)
	if err != nil {
		c.Log.Fatal(err)
	}

	return &messenger{
		Config:  c,
		adapter: a,
	}
}

// Will send a message to the queue, this will be in JSON format.
// The message needs to support JSON marshalling.
//
// The queue name will be prefixed with the environment name.
func (m messenger) Dispatch(msg Message) error {
	m.Log.Infow("Dispatching message", "message", msg)

	json, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	err = m.adapter.Dispatch(adapterMessage{
		Queue:      m.prefixQueue(msg.Queue()),
		Identifier: msg.Identifier(),
		Body:       string(json),
	})
	if err != nil {
		m.Log.Errorw("Error dispatching message", "message", msg, "error", err)
	} else {
		m.Log.Infow("Message dispatched", "message", msg)
	}

	return err
}

// Subscribes to a queue and will handle the messages using the provided handlers.
// All handlers must subscribe to the same queue.
//
// The queue name will be prefixed with the environment name.
//
// This function will block until the shutdown context is cancelled.
//
// If the RestartTimeout is set, the function will restart the subscription upon error.
func (m messenger) Subscribe(h ...MessageHandler) error {
	var queue string
	for _, handler := range h {
		if queue == "" {
			queue = handler.Message().Queue()
		} else if queue != handler.Message().Queue() {
			return ErrDifferentQueues
		}
	}

	queue = m.prefixQueue(queue)
	m.Log.Infof("Subscribing to %s", queue)

	ctx, _ := m.Shutdown.Add()
	defer m.Shutdown.Done()

	// The handleMessage function will be called for each message received from the queue.
	// It will find the correct handler based on the identifier for the message.
	handleMessage := func(a adapterMessage) error {
		for _, handler := range h {
			if a.Identifier == handler.Message().Identifier() {
				msg := handler.Message()
				if err := json.Unmarshal([]byte(a.Body), msg); err != nil {
					m.Log.Error(err)
					return err
				}
				err := handler.Handle(msg)
				if err != nil {
					m.Log.Error(err)
				} else {
					m.Log.Infof("Message %s handled", a.Identifier)
				}
				return err
			}
		}

		err := errors.New("no handler found for message " + a.Identifier)
		m.Log.Error(err.Error())
		return err
	}

	err := m.adapter.Subscribe(queue, handleMessage, ctx)

	if err == nil || err == ctx.Err() {
		return nil
	}

	m.Log.Errorw("Error subscribing to queue", "queue", queue, "error", err)

	if m.RestartTimeout == 0 {
		return err
	}

	m.Log.Infof("Restarting subscription in %s", m.RestartTimeout)
	time.Sleep(m.RestartTimeout)
	return m.Subscribe(h...)
}

// Prefixes the queue name with the environment name.
// This is to prevent queues from different environments from interfering with each other
// when using the same Pub/Sub instance.
func (m messenger) prefixQueue(queue string) string {
	return m.Environment + "." + queue
}
