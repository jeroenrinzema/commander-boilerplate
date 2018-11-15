package rest

import (
	"net/http"
)

// Use use the given middleware on this single route.
// The first argument is the final handle function.
// All the following arguments are middleware handles.
func Use(h http.HandlerFunc, middleware ...func(http.HandlerFunc) http.HandlerFunc) http.HandlerFunc {
	for _, m := range middleware {
		h = m(h)
	}

	return h
}
