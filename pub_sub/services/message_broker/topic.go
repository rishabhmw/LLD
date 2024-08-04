package message_broker

import (
	"fmt"
	"lld/pub_sub/models"
)

type topic struct {
	name               string
	subscriberChannels *models.ConcurrentMap
}

func newTopic(name string) *topic {
	return &topic{
		name:               name,
		subscriberChannels: models.NewConcurrentMap(),
	}
}

func (t *topic) addSubscriberChannel(id string, subscriberChannel messageChannel) error {
	var err error
	t.subscriberChannels.Do(func() {
		_, ok := t.subscriberChannels.GetUnsafe(id)
		if ok {
			err = fmt.Errorf("already subscribed")
			return
		}
		t.subscriberChannels.PutUnsafe(id, subscriberChannel)
	})
	return err
}

func (t *topic) removeSubscriberChannel(id string) {
	t.subscriberChannels.Remove(id)
	fmt.Printf("removed channel %s from subscription of topic %s\n", id, t.name)
}

func (t *topic) publish(message models.Message) {
	t.subscriberChannels.Do(func() {
		ls := t.subscriberChannels.ListUnsafe()
		for _, entry := range ls {
			go func() {
				entry[1].(messageChannel) <- message
			}()
		}
	})
}
