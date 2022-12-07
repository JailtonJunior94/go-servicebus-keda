package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jailtonjunior94/go-servicebus-keda/pkg/bus"
)

func main() {
	ctx := context.TODO()

	serviceBus := bus.NewServiceBus(os.Getenv("CONNECTION_STRING_SB"))
	defer serviceBus.Client.Close(ctx)

	messages := make(chan []byte)

	go serviceBus.ConsumerFromQueue(ctx, "keda-poc-queue", 100, messages)
	for message := range messages {
		fmt.Println(string(message))
	}
}
