package route

import (
	"api/src/middlewares"
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
	routes = append(routes, routeLogin)
	routes = append(routes, routePublications...)

	for _, route := range routes {

		if route.RequireAuthentication {
			r.HandleFunc(route.URI,
				middlewares.Logger(
					middlewares.Authenticate(route.Function),
				),
			).Methods(route.Method)
		} else {
			r.HandleFunc(route.URI,
				middlewares.Logger(route.Function),
			).Methods(route.Method)
		}
	}

	return r
}
