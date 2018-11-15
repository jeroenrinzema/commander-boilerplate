package controllers

import (
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jeroenrinzema/commander"
	"github.com/jeroenrinzema/commander-boilerplate/command/rest"
	uuid "github.com/satori/go.uuid"
)

// OnCommand handles a new command request.
// The received command can be executed in a sync or async manner.
func OnCommand(group *commander.Group) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		res := rest.Response{ResponseWriter: w}
		params := r.URL.Query()
		vars := mux.Vars(r)

		sync, _ := strconv.ParseBool(params.Get("sync"))
		body, _ := ioutil.ReadAll(r.Body)

		action := vars["command"]

		key, KeyErr := uuid.FromString(r.Header.Get("Key"))
		if KeyErr != nil {
			key = uuid.NewV4()
		}

		command := group.NewCommand(action, key, body)

		if sync {
			event, err := group.SyncCommand(command)

			if err != nil {
				res.SendPanic(err.Error(), command)
				return
			}

			if !event.Acknowledged {
				res.SendPanic(event.Action, event)
				return
			}

			res.SendOK(event)
			return
		}

		err := group.AsyncCommand(command)

		if err != nil {
			res.SendPanic(err.Error(), nil)
			return
		}

		res.SendCreated(command)
	}
}
