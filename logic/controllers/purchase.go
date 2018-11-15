package controllers

import (
	"encoding/json"

	"github.com/jeroenrinzema/commander"
	"github.com/jeroenrinzema/commander-boilerplate/logic/models"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// PurchaseModal used during a "Purchase" command
type PurchaseModal struct {
	ID *uuid.UUID `json:"id"`
}

// PurchaseCommand handles the "create" command
type PurchaseCommand struct {
	Database *gorm.DB
}

// Handle handles a "Purchase" command
func (service *PurchaseCommand) Handle(writer commander.ResponseWriter, command *commander.Command) {
	var data PurchaseModal

	UnmarshalErr := json.Unmarshal(command.Data, &data)
	if UnmarshalErr != nil {
		writer.ProduceError("DataParseError", nil)
		return
	}

	cart := models.CartModal{
		ID: data.ID,
	}

	query := service.Database.First(&cart).Update("purchased", true)
	if query.Error != nil {
		writer.ProduceError("CartNotFound", nil)
		return
	}

	res, MarshalErr := json.Marshal(data)
	if MarshalErr != nil {
		writer.ProduceError("ResponseParseError", nil)
		return
	}

	writer.ProduceEvent("Purchased", 1, command.Key, res)
}
