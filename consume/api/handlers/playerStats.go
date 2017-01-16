package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tsap-laval/tsap-common/common/models"
)

// PlayerStatsHandler handles the request to get player stats
func PlayerStatsHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	player := vars["playerId"]

	fmt.Println(player)

	out, err := json.Marshal(models.playerStats{})

	if err != nil {
		panic(err)
	}

	w.Write(out)
}
