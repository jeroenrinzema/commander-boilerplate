package controllers

import (
	"encoding/json"

	"github.com/jeroenrinzema/commander"
	"github.com/jeroenrinzema/commander-boilerplate/projector/models"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// PurchasedModal used during a "Purchased" command
type PurchasedModal struct {
	ID *uuid.UUID `json:"id"`
}

// CartPurchased handles a "Purchased" event
type CartPurchased struct {
	Database *gorm.DB
}

// Handle handles a "Purchased" event
func (service *CartPurchased) Handle(event *commander.Event) {
	var req PurchasedModal

	dataParseError := json.Unmarshal(event.Data, &req)
	if dataParseError != nil {
		return
	}

	cart := models.CartModal{
		ID: req.ID,
	}

	query := service.Database.First(&cart).Update("purchased", true)
	if query.Error != nil {
		return
	}
}
