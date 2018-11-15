package main

import (
	"net/http"

	"github.com/jeroenrinzema/commander-boilerplate/query/common"
	"github.com/jeroenrinzema/commander-boilerplate/query/controllers"
	"github.com/jeroenrinzema/commander-boilerplate/query/rest"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	common.OpenDatabase()
	router := common.OpenWebHub()

	router.HandleFunc("/find/{id}", rest.Use(controllers.FindByID, Authentication)).Methods("GET")
	router.HandleFunc("/find/", rest.Use(controllers.FindAll, Authentication)).Methods("GET")

	http.ListenAndServe(":7070", router)
}

// Authentication validates if the given request is authenticated.
// If the request is not authenticated is a 401 returned.
func Authentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// <- authenticate request and provide the users dataset key
		next.ServeHTTP(w, r)
	}
}
