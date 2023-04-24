package route

import (
	"github.com/gorilla/mux"
	"net/http"
)

// Route represents all routes from API
type Route struct {
	URI                   string
	Method                string
	Function              func(http.ResponseWriter, *http.Request)
	RequireAuthentication bool
}

// Configure configures all routes from API
func Configure(r *mux.Router) *mux.Router {
	routes := routesUsers

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

	return r
}
