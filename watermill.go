package main

import (
	"context"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
	"log"
	"time"
)

// Go 每日一库之 watermill watermill是 Go 语言的一个异步消息解决方案，它支持消息重传、保存消息，后启动的订阅者也能收到前面发布的消息
func main() {
	pubSub := gochannel.NewGoChannel(
		gochannel.Config{},
		watermill.NewStdLogger(false, false),
	)

	messages, err := pubSub.Subscribe(context.Background(), "example.topic")
	if err != nil {
		panic(err)
	}

	go process(messages)

	publishMessage(pubSub)
}

func publishMessage(publisher message.Publisher) {
	for {
		msg := message.NewMessage(watermill.NewUUID(), []byte("Hello World!"))

		if err := publisher.Publish("example.topic", msg); err != nil {
			panic(err)
		}

		time.Sleep(time.Second)
	}
}

func process(messages <-chan *message.Message) {
	for msg := range messages {
		log.Printf("received message: %s, payload: %s", msg.UUID, string(msg.Payload))
		msg.Ack()
	}
}
