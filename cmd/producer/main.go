package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	_ "github.com/jailtonjunior94/go-servicebus-keda/docs"
	"github.com/jailtonjunior94/go-servicebus-keda/pkg/bus"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	swagger "github.com/swaggo/http-swagger"
)

// @title          Producer API
// @version        1.0
// @description    Producer API
// @termsOfService http://swagger.io/terms

// @contact.name  Jailton Junior
// @contact.url   http://jailton.junior.net
// @contact.email jailton.junior94@outlook.com

// @license.name Jailton Junior License
// @license.url  http://jailton.junior.net

// @BasePath /
func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Heartbeat("/health"))

	router.Post("/messages", SendMessageHandle)
	router.Post("/messages-batch", SendMessageBatchHandle)

	router.Get("/docs/*", swagger.Handler(swagger.URL("http://localhost:8000/docs/doc.json")))
	fmt.Printf("🚀 API is running on http://localhost:%v", "8000")
	http.ListenAndServe(":8000", router)

}

type SendMessage struct {
	Message string `json:"message"`
}

// Send Message godoc
// @Summary     Send Message
// @Description Send Message
// @Tags        Messages
// @Accept      json
// @Produce     json
// @Param       request body     SendMessage true "send message request"
// @Success     200     {object} SendMessage
// @Failure     404     {object} SendMessage
// @Failure     500     {object} SendMessage
// @Router      /messages [post]
func SendMessageHandle(w http.ResponseWriter, r *http.Request) {
	var sendMessage SendMessage
	err := json.NewDecoder(r.Body).Decode(&sendMessage)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	ctx, cancel := context.WithTimeout(context.TODO(), 60*time.Second)
	defer cancel()

	serviceBus := bus.NewServiceBus()
	defer serviceBus.Client.Close(ctx)

	err = serviceBus.Publish(ctx, "keda-poc-queue", sendMessage.Message)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Contenty-Type", "application/json")
	json.NewEncoder(w).Encode(SendMessage{Message: "Send message successfully"})
	w.WriteHeader(http.StatusAccepted)
}

type SendMessageBatch struct {
	Messages []string `json:"messages"`
}

// Send Message godoc
// @Summary     Send Message
// @Description Send Message
// @Tags        Messages
// @Accept      json
// @Produce     json
// @Param       request body     SendMessageBatch true "send message request"
// @Success     200     {object} SendMessageBatch
// @Failure     404     {object} SendMessageBatch
// @Failure     500     {object} SendMessageBatch
// @Router      /messages-batch [post]
func SendMessageBatchHandle(w http.ResponseWriter, r *http.Request) {
	var sendMessage SendMessageBatch
	err := json.NewDecoder(r.Body).Decode(&sendMessage)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	ctx, cancel := context.WithTimeout(context.TODO(), 60*time.Second)
	defer cancel()

	serviceBus := bus.NewServiceBus()
	defer serviceBus.Client.Close(ctx)

	err = serviceBus.PublishBatch(ctx, "keda-poc-queue", sendMessage.Messages)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Contenty-Type", "application/json")
	json.NewEncoder(w).Encode(SendMessage{Message: "Send message successfully"})
	w.WriteHeader(http.StatusAccepted)
}
