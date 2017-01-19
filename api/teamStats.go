package api

import (
	"fmt"
	"net/http"
	"strconv"

	"encoding/json"

	"github.com/TSAP-Laval/consume-backend/stats"
	"github.com/gorilla/mux"
)

// TeamStatsHandler handles the request to get team stats
func (c *ConsumeService) TeamStatsHandler(w http.ResponseWriter, r *http.Request) {

	var errMsg string

	vars := mux.Vars(r)

	team := vars["teamID"]

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

	teamStats, err := stats.GetTeamStats(uint(teamID), season.ID, c.datasource)

	if err != nil {
		c.Error(fmt.Sprintf("Error fetching teamstats: %s", err))
		w.WriteHeader(http.StatusNotFound)

		c.ErrorWrite(err.Error(), w)
		return
	}

	bytes, err := json.Marshal(teamStats)

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
