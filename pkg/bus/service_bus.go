package bus

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
)

type ServiceBus struct {
	Client *azservicebus.Client
}

func NewServiceBus() *ServiceBus {
	connectionString := ""
	client, err := azservicebus.NewClientFromConnectionString(connectionString, nil)

	if err != nil {
		panic(err)
	}

	return &ServiceBus{Client: client}
}

func (s *ServiceBus) Publish(ctx context.Context, queueOrTopic, message string) error {
	sender, err := s.Client.NewSender(queueOrTopic, nil)
	if err != nil {
		return err
	}
	defer sender.Close(ctx)

	return sender.SendMessage(ctx, &azservicebus.Message{
		Body: []byte(message),
	}, nil)
}

func (s *ServiceBus) PublishBatch(ctx context.Context, queueOrTopic string, messages []string) error {
	sender, err := s.Client.NewSender(queueOrTopic, nil)
	if err != nil {
		return err
	}
	defer sender.Close(ctx)

	batch, err := sender.NewMessageBatch(ctx, nil)
	if err != nil {
		return err
	}

	for _, message := range messages {
		batch.AddMessage(&azservicebus.Message{
			Body: []byte(message),
		}, nil)
	}

	return sender.SendMessageBatch(ctx, batch, nil)
}

func (s *ServiceBus) ConsumerFromQueue(ctx context.Context, queueName string, count int, out chan<- []byte) {
	receiver, err := s.Client.NewReceiverForQueue(queueName, nil)
	if err != nil {
		panic(err)
	}
	defer receiver.Close(ctx)

	for {
		messages, err := receiver.ReceiveMessages(ctx, count, nil)
		if err != nil {
			panic(err)
		}

		for _, message := range messages {
			out <- message.Body
			if err := receiver.CompleteMessage(ctx, message, nil); err != nil {
				panic(err)
			}
		}
	}
}
