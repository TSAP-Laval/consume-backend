package core

import "net/http"

// Route représente une simple Route gérée par l'API
type Route struct {
	Method  string
	Path    string
	Handler http.HandlerFunc
}
