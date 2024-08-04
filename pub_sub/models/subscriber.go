package models

import (
	"context"
)

type OnNotifyFunc func(m Message)

type ISubscriber interface {
	GetID(ctx context.Context) string
	Listen(ctx context.Context)
	GetChan() chan Message
}

type BaseSubscriber struct {
	subscriberID string
	channel      chan Message
	onNotifyFunc OnNotifyFunc
}

func NewBaseSubscriber(ctx context.Context, subscriberID string, f OnNotifyFunc) ISubscriber {
	return &BaseSubscriber{
		subscriberID: subscriberID,
		channel:      make(chan Message, 10),
		onNotifyFunc: f,
	}
}

func (b *BaseSubscriber) GetID(ctx context.Context) string {
	return b.subscriberID
}

func (b *BaseSubscriber) Listen(ctx context.Context) {
	for b.channel != nil {
		m := <-b.channel
		b.onNotifyFunc(m)
	}
}

func (b *BaseSubscriber) GetChan() chan Message {
	return b.channel
}
