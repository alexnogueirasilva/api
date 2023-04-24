package route

import "net/http"

// Route represents all routes from API
type Route struct {
	URI                   string
	Method                string
	Function              func(http.ResponseWriter, *http.Request)
	RequireAuthentication bool
}
