package teamsmodule

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/TSAP-Laval/common"
	"github.com/TSAP-Laval/consume-backend/app/core"
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
