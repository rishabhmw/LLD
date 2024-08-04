package main

import (
	"context"
	"fmt"
	"lld/pub_sub/models"
	"lld/pub_sub/services/message_broker"
	"time"
)

func main() {
	ctx := context.Background()
	broker := message_broker.NewBaseMessageBroker(ctx)

	if err := broker.CreateTopic("click"); err != nil {
		println("error creating topic: %v", err)
	}
	if err := broker.CreateTopic("hover"); err != nil {
		println("error creating topic: %v", err)
	}

	s := models.NewBaseSubscriber(ctx, "telem-service", func(m models.Message) {
		fmt.Printf("reaching out to telem-service about action: %s\n", string(m.Data))
	})
	s2 := models.NewBaseSubscriber(ctx, "ad-service", func(m models.Message) {
		fmt.Printf("triggering ad-service about action: %s\n", string(m.Data))
	})

	if err := broker.SubscribeToTopic("click", s); err != nil {
		println("error subscribing to topic: %v", err)
	}

	if err := broker.SubscribeToTopic("click", s2); err != nil {
		println("error subscribing to topic: %v", err)
	}
	if err := broker.SubscribeToTopic("hover", s); err != nil {
		println("error subscribing to topic: %v", err)
	}

	go s.Listen(ctx)
	go s2.Listen(ctx)

	go func() {
		if err := broker.Publish(ctx, models.Message{
			ID:    "m1",
			Data:  []byte("button click"),
			Topic: "click",
		}); err != nil {
			println("error publishing to topic: %v", err)
		}
	}()
	go func() {
		if err := broker.Publish(ctx, models.Message{
			ID:    "m1",
			Data:  []byte("menu click"),
			Topic: "click",
		}); err != nil {
			println("error publishing to topic: %v", err)
		}
	}()
	go func() {
		if err := broker.UnsubscribeFromTopic("click", s); err != nil {
			println("error unsubscribing from topic: %v", err)
		}
	}()
	go func() {
		if err := broker.Publish(ctx, models.Message{
			ID:    "m1",
			Data:  []byte("download hover"),
			Topic: "hover",
		}); err != nil {
			println("error publishing to topic: %v", err)
		}
	}()
	go func() {
		if err := broker.Publish(ctx, models.Message{
			ID:    "m1",
			Data:  []byte("download click"),
			Topic: "click",
		}); err != nil {
			println("error publishing to topic: %v", err)
		}
	}()
	go func() {
		if err := broker.Publish(ctx, models.Message{
			ID:    "m1",
			Data:  []byte("profile click"),
			Topic: "click",
		}); err != nil {
			println("error publishing to topic: %v", err)
		}
	}()
	go func() {
		if err := broker.Publish(ctx, models.Message{
			ID:    "m1",
			Data:  []byte("profile hover"),
			Topic: "click",
		}); err != nil {
			println("error publishing to topic: %v", err)
		}
	}()
	time.Sleep(5 * time.Second)
	println("<<<end>>>")
}
