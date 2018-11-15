package groups

import (
	"time"

	"github.com/jeroenrinzema/commander"
)

// Cart ...
var Cart = &commander.Group{
	Timeout: 5 * time.Second,
	Topics: []commander.Topic{
		commander.Topic{
			Name:    "cart-events",
			Type:    commander.EventTopic,
			Consume: true,
			Produce: false,
		},
		commander.Topic{
			Name:    "cart-commands",
			Type:    commander.CommandTopic,
			Consume: false,
			Produce: true,
		},
	},
}
