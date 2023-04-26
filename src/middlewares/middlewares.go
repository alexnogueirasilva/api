package middlewares

import (
	"api/src/authentication"
	"log"
	"net/http"
)

// Logger writes information about the requests on the terminal
func Logger(nextFunction http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n%s %s %s", r.Method, r.RequestURI, r.Host)
		nextFunction(w, r)
	}

}

// Authenticate verifies if the user making the request is authenticated
func Authenticate(nextFunction http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := authentication.ValidateToken(r); err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		nextFunction(w, r)
	}
}
