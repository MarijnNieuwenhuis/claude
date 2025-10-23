package messenger

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"sync"
	"time"

	"cloud.google.com/go/pubsub"
	"go.uber.org/zap"
)

type PubsubConfig struct {
	Emulator        string
	Project         string
	DeadLetterTopic string
}

type pubsubAdapter struct {
	config PubsubConfig
	client *pubsub.Client
	topics map[string]*pubsub.Topic
	log    *zap.SugaredLogger
	sync.Mutex
}

type pubsubMessage struct {
	Headers pubsubHeaders `json:"headers"`
	Body    string        `json:"body"`
}

type pubsubHeaders struct {
	Type string `json:"type"`
}

var ErrMissingProject = errors.New("missing project")

// The creation of the adapter will create a new Pub/Sub client using the provided configuration.
func newPubsubAdapter(c PubsubConfig, log *zap.SugaredLogger) (*pubsubAdapter, error) {
	if c.Emulator != "" {
		os.Setenv("PUBSUB_EMULATOR_HOST", c.Emulator)
		if c.Project == "" {
			c.Project = "emulator-project"
		}
	}

	if c.Project == "" {
		return nil, ErrMissingProject
	}

	client, err := pubsub.NewClient(context.Background(), c.Project)
	if err != nil {
		return nil, err
	}

	return &pubsubAdapter{
		config: c,
		client: client,
		topics: make(map[string]*pubsub.Topic),
		log:    log,
	}, nil
}

// Dispatch will send a message to the queue, this will be in JSON format.
// The message needs to support JSON marshalling.
//
// This method assumes that the topic already exists.
func (p *pubsubAdapter) Dispatch(msg adapterMessage) error {
	m := pubsubMessage{
		Headers: pubsubHeaders{
			Type: msg.Identifier,
		},
		Body: string(msg.Body),
	}
	json, err := json.Marshal(m)
	if err != nil {
		return err
	}

	topic, err := p.topic(msg.Queue, false)
	if err != nil {
		return err
	}

	res := topic.Publish(context.Background(), &pubsub.Message{
		Data: json,
	})
	_, err = res.Get(context.Background())
	return err
}

// Subscribe will listen to the queue and call the provided handler when a message is received.
// This is a blocking method and will return when the context is cancelled.
//
// If the subscription and/or topic do not exist, they will be created.
// If they do exist, they will be updated to make sure they are correctly configured to prevent
// alterations in the Google console.
func (p *pubsubAdapter) Subscribe(queue string, h handleMessage, ctx context.Context) error {
	sub, _, err := p.subscription(queue, queue, p.config.DeadLetterTopic)
	if err != nil {
		return err
	}

	p.log.Infof("Listening to Pub/Sub subscription %s", sub.ID())

	return sub.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		p.log.Infow("Received Pub/Sub message", "id", msg.ID, "queue", queue, "data", string(msg.Data))

		var m pubsubMessage
		if err := json.Unmarshal(msg.Data, &m); err != nil {
			msg.Nack()
			return
		}

		if err := h(adapterMessage{
			Queue:      queue,
			Identifier: m.Headers.Type,
			Body:       m.Body,
		}); err != nil {
			msg.Nack()
			return
		}

		msg.Ack()
	})
}

// Retrieve the topic and create it if it does not exist.
func (p *pubsubAdapter) topic(queue string, create bool) (*pubsub.Topic, error) {
	if topic, ok := p.topics[queue]; ok {
		return topic, nil
	}

	topic := p.client.Topic(queue)
	if create {
		err := p.createTopicIfNotExists(topic)
		if err != nil {
			return nil, err
		}
	}

	p.Lock()
	defer p.Unlock()
	p.topics[queue] = topic

	return topic, nil
}

func (p *pubsubAdapter) createTopicIfNotExists(topic *pubsub.Topic) error {
	if exists, err := topic.Exists(context.Background()); exists || err != nil {
		return err
	}

	p.log.Infof("Creating Pub/Sub topic %s", topic.ID())
	_, err := p.client.CreateTopic(context.Background(), topic.ID())

	return err
}

// Retrieve the subscription and create it if it does not exist.
// The subscription will be updated to make sure it is correctly configured.
//
// This method will also make sure the dead letter topic and subscription are correctly configured.
func (p *pubsubAdapter) subscription(subscription, topic, deadLetterTopic string) (*pubsub.Subscription, *pubsub.Topic, error) {
	top, err := p.topic(topic, true)
	if err != nil {
		return nil, nil, err
	}

	sub := p.client.Subscription(subscription)
	p.createSubscriptionIfNotExists(sub, top)

	if deadLetterTopic == "" {
		return sub, top, nil
	}

	// Make sure the dead letter topic & subscription exists.
	_, dlTop, err := p.subscription(deadLetterTopic, deadLetterTopic, "")
	if err != nil {
		return nil, nil, err
	}

	p.log.Infof("Updating Pub/Sub subscription %s", subscription)
	_, err = sub.Update(context.Background(), pubsub.SubscriptionConfigToUpdate{
		DeadLetterPolicy: &pubsub.DeadLetterPolicy{
			DeadLetterTopic:     dlTop.String(),
			MaxDeliveryAttempts: 5,
		},
		RetryPolicy: &pubsub.RetryPolicy{
			MinimumBackoff: 10 * time.Second,
			MaximumBackoff: 300 * time.Second,
		},
	})

	return sub, top, err
}

func (p *pubsubAdapter) createSubscriptionIfNotExists(sub *pubsub.Subscription, topic *pubsub.Topic) error {
	if exists, err := sub.Exists(context.Background()); exists || err != nil {
		return err
	}

	p.log.Infof("Creating Pub/Sub subscription %s", sub.ID())
	_, err := p.client.CreateSubscription(context.Background(), sub.ID(), pubsub.SubscriptionConfig{
		Topic: topic,
	})

	return err
}
