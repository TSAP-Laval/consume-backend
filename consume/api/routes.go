package api

import "github.com/gorilla/mux"

// GetRouter  Definit les routes de l'app.
func GetRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", IndexHandler)
	return r
}
