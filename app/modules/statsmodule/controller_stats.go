package statsmodule

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/TSAP-Laval/common"
	"github.com/TSAP-Laval/consume-backend/app/core"
	"github.com/TSAP-Laval/consume-backend/stats"
	"github.com/gorilla/mux"
)

// StatsController répond aux routes de statistiques
type StatsController struct {
	core.Controller
	datasource common.IDatasource
	config     *core.ConsumeConfiguration
	logger     *log.Logger
}

// NewStatsController instancie un nouveau controlleur
func NewStatsController(datasource common.IDatasource, config *core.ConsumeConfiguration, logger *log.Logger) *StatsController {
	return &StatsController{
		Controller: core.Controller{},
		datasource: datasource,
		config:     config,
		logger:     logger,
	}
}

// GetPlayerStats handles the request to get player stats
func (c *StatsController) GetPlayerStats(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	player := vars["playerID"]
	team := vars["teamID"]

	playerID, err := strconv.Atoi(player)

	if err != nil {
		c.SendJSON(w, core.ErrorMessage{
			Error: fmt.Sprintf("PlayerID %s invalid", player),
		}, http.StatusBadRequest)
		return
	}

	teamID, err := strconv.Atoi(team)

	if err != nil {
		c.SendJSON(w, core.ErrorMessage{
			Error: fmt.Sprintf("TeamID %s invalid", team),
		}, http.StatusBadRequest)
		return
	}

	seasonIDRaw := r.URL.Query().Get("season")

	var seasonID uint

	if seasonIDRaw != "" {
		seasonIDu, err := strconv.Atoi(seasonIDRaw)

		if err != nil {
			c.SendJSON(w, core.ErrorMessage{
				Error: fmt.Sprintf("Season %s invalid", seasonIDRaw),
			}, http.StatusBadRequest)
		}

		seasonID = uint(seasonIDu)

	} else {
		season, err := c.datasource.GetCurrentSeason()

		if c.HandleError(err, w) {
			return
		}

		seasonID = season.ID
	}

	positionIDRaw := r.URL.Query().Get("position")

	var positionID uint

	if positionIDRaw != "" {
		positionIDu, err := strconv.Atoi(positionIDRaw)

		if err != nil {
			c.SendJSON(w, core.ErrorMessage{
				Error: fmt.Sprintf("Position %s invalid", positionIDRaw),
			}, http.StatusBadRequest)
		}

		positionID = uint(positionIDu)
	}

	stats, err := stats.GetPlayerStats(uint(playerID), uint(teamID), seasonID, positionID, c.datasource)

	if err != nil {
		c.SendJSON(w, core.ErrorMessage{
			Error: fmt.Sprintf("Error fetching stats: %s", err),
		}, http.StatusNotFound)
		return
	}

	c.SendJSON(w, stats, http.StatusOK)
}

// GetPlayerMatchStats handles the request to get player stats per match
func (c *StatsController) GetPlayerMatchStats(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	player := vars["playerID"]
	match := vars["matchID"]

	playerID, err := strconv.Atoi(player)

	if err != nil {
		c.SendJSON(w, core.ErrorMessage{
			Error: fmt.Sprintf("PlayerID %s invalid", player),
		}, http.StatusBadRequest)
		return
	}

	matchID, err := strconv.Atoi(match)
	if err != nil {
		c.SendJSON(w, core.ErrorMessage{
			Error: fmt.Sprintf("MatchID %s invalid", match),
		}, http.StatusBadRequest)
		return
	}

	stats, err := stats.GetPlayerActions(uint(playerID), uint(matchID), c.datasource)

	if err != nil {
		c.SendJSON(w, core.ErrorMessage{
			Error: fmt.Sprintf("Error fetching stats: %s", err),
		}, http.StatusBadRequest)
		return
	}

	c.SendJSON(w, stats, http.StatusOK)
}

// GetTeamStats handles the request to get team stats
func (c *StatsController) GetTeamStats(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	team := vars["teamID"]

	teamID, err := strconv.Atoi(team)

	if err != nil {
		c.SendJSON(w, core.ErrorMessage{
			Error: fmt.Sprintf("TeamID %s invalid", team),
		}, http.StatusBadRequest)
		return
	}

	season, err := c.datasource.GetCurrentSeason()

	if c.HandleError(err, w) {
		return
	}

	teamStats, err := stats.GetTeamStats(uint(teamID), season.ID, c.datasource)

	if err != nil {
		c.SendJSON(w, core.ErrorMessage{
			Error: fmt.Sprintf("Error fetching teamstats: %s", err),
		}, http.StatusNotFound)
		return
	}

	c.SendJSON(w, teamStats, http.StatusOK)
}

// GetPlayerPositions retourne la liste de positions occupées par le joueur
func (c *StatsController) GetPlayerPositions(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	player := vars["playerID"]

	playerID, err := strconv.Atoi(player)

	if err != nil {
		c.SendJSON(w, core.ErrorMessage{
			Error: fmt.Sprintf("PlayerID %s invalid", player),
		}, http.StatusBadRequest)
		return
	}

	positions, err := c.datasource.GetPositions(uint(playerID))

	if c.HandleError(err, w) {
		return
	}

	c.SendJSON(w, positions, http.StatusOK)
}
