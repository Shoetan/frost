package handlers

import (
	"net/http"
	"io"
)

func Request() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		content, err := io.ReadAll(r.Body)

		if err != nil {
			http.Error(w, "Could not read contents from body", http.StatusBadRequest)
			return
		}
		contentType := r.Header.Get("Content-Type")
    w.Header().Set("Content-Type", contentType)

		_, err = w.Write(content)

		if err != nil {
			http.Error(w, "Could not write content", http.StatusBadRequest)
		}
		
	}
	
}