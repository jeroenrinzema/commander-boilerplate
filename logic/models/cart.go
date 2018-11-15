package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// CartModal holds the structure of the carts table
type CartModal struct {
	ID        *uuid.UUID `gorm:"type:uuid; primary_key"`
	Customer  string     `gorm:"not null"`
	Purchased bool       `gorm:"default:'false'"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

// TableName returns the table name of CartModal
func (modal *CartModal) TableName() string {
	return "ServiceCartsView"
}
