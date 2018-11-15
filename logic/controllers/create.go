package controllers

import (
	"encoding/json"

	"github.com/jeroenrinzema/commander"
	"github.com/jeroenrinzema/commander-boilerplate/logic/models"
	"github.com/jinzhu/gorm"
)

// CreateCommand handles the "create" command
type CreateCommand struct {
	Database *gorm.DB
}

// CreateModal is used during a "create" command
type CreateModal struct {
	Customer string `json:"customer"`
}

// Handle handles a "create" command and initializes a new shopping cart
func (service *CreateCommand) Handle(writer commander.ResponseWriter, command *commander.Command) {
	var data CreateModal

	// Parse the data back to a struct
	UnmarshalErr := json.Unmarshal(command.Data, &data)
	if UnmarshalErr != nil {
		writer.ProduceError("DataParseError", nil)
		return
	}

	cart := models.CartModal{
		ID:       &command.Key,
		Customer: data.Customer,
	}

	query := service.Database.Create(&cart)
	if query.Error != nil {
		writer.ProduceError("ModelError", nil)
		return
	}

	res, MarshalErr := json.Marshal(cart)
	if MarshalErr != nil {
		writer.ProduceError("ResponseParseError", nil)
		return
	}

	writer.ProduceEvent("Created", 1, command.Key, res)
}
