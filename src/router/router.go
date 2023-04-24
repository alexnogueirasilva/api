package router

import (
	"api/src/router/route"
	"github.com/gorilla/mux"
)

// Generate returns a new router instance
func Generate() *mux.Router {
	r := mux.NewRouter()
	return route.Configure(r)
}
