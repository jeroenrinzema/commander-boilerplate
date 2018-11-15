package controllers

import (
	"encoding/json"

	"github.com/jeroenrinzema/commander"
	"github.com/jeroenrinzema/commander-boilerplate/projector/models"
	"github.com/jinzhu/gorm"
)

// CreateCart handles a "created" event
type CreateCart struct {
	Database *gorm.DB
}

// Handle handles a "created" event
func (service *CreateCart) Handle(event *commander.Event) {
	var cart models.CartModal

	dataParseError := json.Unmarshal(event.Data, &cart)
	// Parse the data back to a struct
	if dataParseError != nil {
		return
	}

	query := service.Database.Create(&cart)
	if query.Error != nil {
		return
	}
}
