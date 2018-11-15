package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jeroenrinzema/commander"
	"github.com/jeroenrinzema/commander-boilerplate/command/websocket"
)

// ConsumeEvents starts consuming events
// all received events are broadcasted over the websocket network
func ConsumeEvents(hub *websocket.Hub, group *commander.Group) {
	events, closing := group.NewEventConsumer()
	defer closing()

	for {
		event := <-events
		data, err := json.Marshal(event)

		if err != nil {
			continue
		}

		hub.Broadcast(string(data))
	}
}

// OnWebsocket handles a new websocket request.
// The session is added into the websocket hub and receives new events.
func OnWebsocket(hub *websocket.Hub, group *commander.Group) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := websocket.Upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println("upgrade:", err)
			return
		}

		client := &websocket.Client{
			Hub:  hub,
			Conn: conn,
			Send: make(chan []byte, 256),
		}

		client.Hub.Register <- client
		go client.WritePump()
	}
}
