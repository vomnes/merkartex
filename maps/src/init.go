package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type adapter func(http.Handler) http.Handler

// adapt transforms an handler without changing it's type. Usefull for authentification.
func adapt(h http.Handler, adapters ...adapter) http.Handler {
	for _, adapter := range adapters {
		h = adapter(h)
	}
	return h
}

// adapt the request by checking the auth and filling the context with usefull data
func enhanceHandlers(r *mux.Router) http.Handler {
	return adapt(r, withRights())
}

// withRights is an adapter that check authentification
func withRights() adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// No authentification to check
			h.ServeHTTP(w, r)
		})
	}
}
