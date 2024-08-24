package handlers

import (
	"net/http"
)

func Request() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Working"))
	}
	
}