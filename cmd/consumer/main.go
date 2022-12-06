package main

import (
	"context"
	"fmt"

	"github.com/jailtonjunior94/go-servicebus-keda/pkg/bus"
)

func main() {
	ctx := context.TODO()

	serviceBus := bus.NewServiceBus()
	defer serviceBus.Client.Close(ctx)

	messages := make(chan []byte)

	go serviceBus.ConsumerFromQueue(ctx, "keda-poc-queue", 100, messages)
	for message := range messages {
		fmt.Println(string(message))
	}
}
