package api

import (
	"fmt"
	"net/http"
	"strconv"

	"encoding/json"

	"github.com/TSAP-Laval/consume-backend/stats"
	"github.com/gorilla/mux"
)

// PlayerMatchStatsHandler handles the request to get player stats
func (c *ConsumeService) PlayerMatchStatsHandler(w http.ResponseWriter, r *http.Request) {

	var errMsg string

	vars := mux.Vars(r)

	player := vars["playerID"]
	match := vars["matchID"]

	playerID, err := strconv.Atoi(player)

	if err != nil {
		c.Error(fmt.Sprintf("PlayerID %s invalid.", player))
		w.WriteHeader(http.StatusBadRequest)
		c.ErrorWrite("PlayerID is invalid", w)
		return
	}

	matchID, err := strconv.Atoi(match)
	if err != nil {
		c.Error(fmt.Sprintf("MatchID %s invalid.", match))
		w.WriteHeader(http.StatusBadRequest)
		c.ErrorWrite("MatchID is invalid", w)
		return
	}

	stats, err := stats.GetPlayerActions(uint(playerID), uint(matchID), c.datasource)

	if err != nil {
		c.Error(fmt.Sprintf("Error fetching stats: %s", err))
		w.WriteHeader(http.StatusNotFound)

		c.ErrorWrite(err.Error(), w)
		return
	}

	bytes, err := json.Marshal(stats)

	if err != nil {
		c.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)

		if c.config.Debug {
			errMsg = err.Error()
		} else {
			errMsg = "An error occured"
		}

		c.ErrorWrite(errMsg, w)
		return
	}

	_, err = w.Write(bytes)

	if err != nil {
		c.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
}
