package api

import (
	"fmt"
	"net/http"
	"strconv"

	"encoding/json"

	"github.com/gorilla/mux"
	"github.com/tsap-laval/consume-backend/stats"
)

// PlayerStatsHandler handles the request to get player stats
func (c *ConsumeService) PlayerStatsHandler(w http.ResponseWriter, r *http.Request) {

	var errMsg string

	vars := mux.Vars(r)

	player := vars["playerID"]
	team := vars["teamID"]

	playerID, err := strconv.Atoi(player)

	if err != nil {
		c.Error(fmt.Sprintf("PlayerID %s invalid.", player))
		w.WriteHeader(http.StatusBadRequest)
		c.ErrorWrite("PlayerID is invalid", w)
		return
	}

	teamID, err := strconv.Atoi(team)
	if err != nil {
		c.Error(fmt.Sprintf("TeamID %s invalid.", team))
		w.WriteHeader(http.StatusBadRequest)
		c.ErrorWrite("TeamID is invalid", w)
		return
	}

	season, err := c.datasource.GetCurrentSeason()

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

	stats, err := stats.GetPlayerStats(uint(playerID), uint(teamID), season.ID, c.datasource)

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
