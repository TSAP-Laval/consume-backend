package teamsmodule

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/TSAP-Laval/common"
	"github.com/TSAP-Laval/consume-backend/app/core"
	"github.com/TSAP-Laval/consume-backend/stats"
	"github.com/TSAP-Laval/models"
	"github.com/gorilla/mux"
)

// TeamsController gère les requêtes de gestion des métriques
type TeamsController struct {
	core.Controller
	datasource common.IDatasource
	config     *core.ConsumeConfiguration
	logger     *log.Logger
}

// NewTeamsController instancie un TeamsController
func NewTeamsController(datasource common.IDatasource, config *core.ConsumeConfiguration, logger *log.Logger) *TeamsController {
	return &TeamsController{
		Controller: core.Controller{},
		datasource: datasource,
		config:     config,
		logger:     logger,
	}
}

// GetMapParameters gère la récupération de la taille de la map
func (c *TeamsController) GetMapParameters(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	teamIDRaw := vars["teamID"]

	teamID, err := strconv.Atoi(teamIDRaw)

	if err != nil {
		c.SendJSON(w, core.ErrorMessage{
			Error: fmt.Sprintf("TeamID %s invalid", teamIDRaw),
		}, http.StatusBadRequest)
	}
	Teams, err := c.datasource.GetMapSize(uint(teamID))

	if c.HandleError(err, w) {
		return
	}

	var displayParams = MapParamsDisplaySchema{
		ID:     Teams.ID,
		Width:  Teams.Longeur,
		Height: Teams.Largeur,
	}

	c.SendJSON(w, displayParams, http.StatusOK)
}

// SetMapParameters modifie la préférence de la map pour une équipe
func (c *TeamsController) SetMapParameters(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	teamIDRaw := vars["teamID"]

	teamID, err := strconv.Atoi(teamIDRaw)

	if err != nil {
		c.SendJSON(w, core.ErrorMessage{
			Error: fmt.Sprintf("TeamID %s invalid", teamIDRaw),
		}, http.StatusBadRequest)
	}

	var params MapParamsCreationSchema
	if err := c.GetContent(&params, r); err != nil {
		return
	}

	if params.Height == 0 || params.Width == 0 {
		c.SendJSON(w, core.ErrorMessage{
			Error: "Invalid Payload",
		}, http.StatusBadRequest)
		return
	}

	err = c.datasource.SetMapSize(params.Width, params.Height, uint(teamID))

	if c.HandleError(err, w) {
		return
	}

	c.SendJSON(w, core.SimpleMessage{
		Body: "ok",
	}, http.StatusCreated)
}

// GetTeamMatches permet d'obtenir tous les matchs d'une équipe
func (c *TeamsController) GetTeamMatches(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	teamIDRaw := vars["teamID"]

	teamID, err := strconv.Atoi(teamIDRaw)

	if err != nil {
		c.SendJSON(w, core.ErrorMessage{
			Error: fmt.Sprintf("TeamID %s invalid", teamIDRaw),
		}, http.StatusBadRequest)
	}

	seasonIDRaw := vars["seasonID"]

	seasonID, err := strconv.Atoi(seasonIDRaw)

	if err != nil {
		c.SendJSON(w, core.ErrorMessage{
			Error: fmt.Sprintf("SeasonID %s invalid", seasonIDRaw),
		}, http.StatusBadRequest)
	}

	Matches, err := c.datasource.GetMatchesInfos(uint(teamID), uint(seasonID))

	if c.HandleError(err, w) {
		return
	}
	displayMatches := make([]MatchesDisplaySchema, len(*Matches))
	for i, m := range *Matches {
		displayMatches[i] = MatchesDisplaySchema{
			ID:   m.ID,
			Lieu: m.Lieu.Nom,
			Date: m.Date,
		}
		if m.EquipeMaisonID == teamID {
			displayMatches[i].EquipeAdverse = m.EquipeAdverse.Nom
			displayMatches[i].Equipe = m.EquipeMaison.Nom
		} else {
			displayMatches[i].EquipeAdverse = m.EquipeMaison.Nom
			displayMatches[i].Equipe = m.EquipeAdverse.Nom
		}
	}

	c.SendJSON(w, displayMatches, http.StatusOK)
}

// GetTeamActions permet d'obtenir les actions de tous les joueurs pour un match
func (c *TeamsController) GetTeamActions(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	teamIDRaw := vars["teamID"]

	teamID, err := strconv.Atoi(teamIDRaw)

	if err != nil {
		c.SendJSON(w, core.ErrorMessage{
			Error: fmt.Sprintf("TeamID %s invalid", teamIDRaw),
		}, http.StatusBadRequest)
	}

	matchIDRaw := vars["matchID"]

	matchID, err := strconv.Atoi(matchIDRaw)

	if err != nil {
		c.SendJSON(w, core.ErrorMessage{
			Error: fmt.Sprintf("MatchID %s invalid", matchIDRaw),
		}, http.StatusBadRequest)
	}
	match, err := c.datasource.GetMatchActions(uint(teamID), uint(matchID))
	var joueurs []models.Joueur
	if match.EquipeMaisonID == teamID {
		joueurs = match.EquipeMaison.Joueurs
	} else if match.EquipeAdverseID == teamID {
		joueurs = match.EquipeAdverse.Joueurs
	} else {
		c.SendJSON(w, core.ErrorMessage{
			Error: fmt.Sprintf("This match doesn't belong to this team"),
		}, http.StatusNotFound)
		return
	}

	matchActions := MatchActions{}
	matchActions.MatchID = uint(matchID)
	matchActions.TeamID = uint(teamID)
	matchActions.Date = match.Date
	players := make([]*stats.PlayerMatchStats, len(joueurs))
	for i, m := range joueurs {
		players[i], err = stats.GetPlayerActions(m.ID, uint(matchID), c.datasource)
		if err != nil {
			c.SendJSON(w, core.ErrorMessage{
				Error: fmt.Sprintf("Error fetching stats: %s", err),
			}, http.StatusBadRequest)
			return
		}
	}
	matchActions.Players = players

	c.SendJSON(w, matchActions, http.StatusOK)
}
