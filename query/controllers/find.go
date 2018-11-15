package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jeroenrinzema/commander-boilerplate/query/common"
	"github.com/jeroenrinzema/commander-boilerplate/query/models"
	"github.com/jeroenrinzema/commander-boilerplate/query/rest"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// FindByID finds a cart by the given id
func FindByID(w http.ResponseWriter, r *http.Request) {
	res := rest.Response{ResponseWriter: w}
	vars := mux.Vars(r)

	id, err := uuid.FromString(vars["id"])

	if err != nil {
		res.SendPanic(err.Error(), nil)
		return
	}

	cart := models.Carts{}
	query := common.Database.Where(models.Carts{ID: &id}).First(&cart)

	if query.Error == gorm.ErrRecordNotFound {
		res.SendNotFound()
		return
	}

	if query.Error != nil {
		res.SendPanic(query.Error.Error(), nil)
		return
	}

	res.SendOK(cart)
}

// FindAll returns all carts found in the database
func FindAll(w http.ResponseWriter, r *http.Request) {
	res := rest.Response{ResponseWriter: w}

	carts := []models.Carts{}
	query := common.Database.Find(&carts)

	if query.Error != nil {
		res.SendPanic(query.Error.Error(), nil)
		return
	}

	res.SendOK(carts)
}
