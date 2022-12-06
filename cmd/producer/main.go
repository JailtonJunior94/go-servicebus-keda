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

	err := serviceBus.Publish(ctx, "keda-poc-queue", "Agora foi")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Send message successfully")
}
