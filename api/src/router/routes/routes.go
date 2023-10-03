package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Uri                    string
	Method                 string
	Function               func(http.ResponseWriter, *http.Request)
	AuthenticationRequired bool
}

func Configure(r *mux.Router) *mux.Router {
	routes := UsersRoutes
	for _, route := range routes {
		r.HandleFunc(route.Uri, route.Function).Methods(route.Method)
	}
	return r
}
