package api

import (
	"github.com/gorilla/mux"
	"github.com/tsap-laval/consume-backend/consume/api/handlers"
)

func getRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/stats/player/{playerId}", handlers.PlayerStatsHandler)
	return r
}
