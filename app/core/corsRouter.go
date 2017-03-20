package core

import "github.com/gorilla/mux"
import "net/http"

// CORSRouter est un routeur s'occupant du Cross Origin Resource Sharing
type CORSRouter struct {
	R *mux.Router
}

func (s *CORSRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if origin := r.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	}

	if r.Method == "OPTIONS" {
		return
	}

	s.R.ServeHTTP(w, r)
}
