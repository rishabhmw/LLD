package message_broker

import (
	"context"
	"fmt"
	"lld/pub_sub/models"
)

type IMessageBroker interface {
	CreateTopic(name string) error
	SubscribeToTopic(name string, subscriber models.ISubscriber) error
	Publish(ctx context.Context, message models.Message) error
	UnsubscribeFromTopic(name string, subscriber models.ISubscriber) error
}

type messageChannel chan models.Message

type BaseMessageBroker struct {
	topics *models.ConcurrentMap
}

// This can be made singleton?

func NewBaseMessageBroker(ctx context.Context) IMessageBroker {
	return &BaseMessageBroker{
		topics: models.NewConcurrentMap(),
	}
}

func (b *BaseMessageBroker) CreateTopic(name string) error {
	var err error
	b.topics.Do(func() {
		if _, found := b.topics.GetUnsafe(name); found {
			err = fmt.Errorf("topic %s already exists", name)
			return
		}
		b.topics.PutUnsafe(name, newTopic(name))
	})

	return err
}

func (b *BaseMessageBroker) SubscribeToTopic(name string, subscriber models.ISubscriber) error {
	var err error
	b.topics.Do(func() {
		t, found := b.topics.GetUnsafe(name)
		if !found {
			err = fmt.Errorf("topic %s does not exist", name)
			return
		}
		err = t.(*topic).addSubscriberChannel(subscriber.GetID(context.Background()), subscriber.GetChan())
	})
	return err
}

func (b *BaseMessageBroker) UnsubscribeFromTopic(name string, subscriber models.ISubscriber) error {
	var err error
	b.topics.Do(func() {
		t, ok := b.topics.GetUnsafe(name)
		if !ok {
			err = fmt.Errorf("topic does not exist")
			return
		}
		t.(*topic).removeSubscriberChannel(subscriber.GetID(context.Background()))
	})
	return err
}

func (b *BaseMessageBroker) Publish(ctx context.Context, message models.Message) error {
	var err error
	b.topics.Do(func() {
		t, found := b.topics.GetUnsafe(message.Topic)
		if !found {
			err = fmt.Errorf("topic %s not found", message.Topic)
			return
		}
		t.(*topic).publish(message)
	})
	return err
}
