package main

import (
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gorilla/mux"
	"github.com/jeroenrinzema/commander"
	"github.com/jeroenrinzema/commander-boilerplate/command/controllers"
	"github.com/jeroenrinzema/commander-boilerplate/command/groups"
	"github.com/jeroenrinzema/commander-boilerplate/command/rest"
	"github.com/jeroenrinzema/commander-boilerplate/command/websocket"
)

func main() {
	router := mux.NewRouter()
	socket := websocket.NewHub()

	config := commander.NewConfig()
	config.Brokers = strings.Split(os.Getenv("KAFKA_BROKERS"), ",")
	config.Group = os.Getenv("KAFKA_GROUP")
	config.Kafka = &kafka.ConfigMap{
		"default.topic.config": kafka.ConfigMap{"auto.offset.reset": "earliest"},
	}

	config.AddGroups(groups.Cart)

	client, err := commander.New(config)
	if err != nil {
		panic(err)
	}

	go client.Consume()
	go controllers.ConsumeEvents(socket, groups.Cart)

	router.HandleFunc("/command/{command}", rest.Use(controllers.OnCommand(groups.Cart), Authentication)).Methods("POST")
	router.HandleFunc("/updates", rest.Use(controllers.OnWebsocket(socket, groups.Cart), Authentication)).Methods("GET")

	server := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  7 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	go func() {
		sigterm := make(chan os.Signal, 1)
		signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)

		<-sigterm

		server.Close()
		client.Close()
	}()

	server.ListenAndServe()
}

// Authentication validates if the given request is authenticated.
// If the request is not authenticated is a 401 returned.
func Authentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// <- authenticate request and provide the users dataset key
		next.ServeHTTP(w, r)
	}
}
