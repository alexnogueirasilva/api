package router

import "github.com/gorilla/mux"

// Generate returns a new router instance
func Generate() *mux.Router {
	return mux.NewRouter()
}
