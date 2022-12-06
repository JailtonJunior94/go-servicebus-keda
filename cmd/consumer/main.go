package main

import (
	"context"
	"fmt"
	"time"

	"github.com/jailtonjunior94/go-servicebus-keda/pkg/bus"
)

func main() {
	ctx, cancel := context.WithTimeout(context.TODO(), 60*time.Second)
	defer cancel()

	serviceBus := bus.NewServiceBus()
	defer serviceBus.Client.Close(ctx)

	messages := make(chan []byte)

	go serviceBus.ConsumerFromQueue(ctx, "keda-poc-queue", 100, messages)
	for message := range messages {
		fmt.Println(string(message))
	}
}
