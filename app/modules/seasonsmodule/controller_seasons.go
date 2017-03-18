package seasonsmodule

import (
	"log"
	"net/http"

	"github.com/TSAP-Laval/common"
	"github.com/TSAP-Laval/consume-backend/app/core"
)

// SeasonsController répond aux routes ayant rapport aux saisons
type SeasonsController struct {
	core.Controller
	datasource common.IDatasource
	config     *core.ConsumeConfiguration
	logger     *log.Logger
}

// NewSeasonsController instancie un nouveau controlleur
func NewSeasonsController(datasource common.IDatasource, config *core.ConsumeConfiguration, logger *log.Logger) *SeasonsController {
	return &SeasonsController{
		Controller: core.Controller{},
		datasource: datasource,
		config:     config,
		logger:     logger,
	}
}

// GetAllSeasons retourne la liste de saisons
func (c *SeasonsController) GetAllSeasons(w http.ResponseWriter, r *http.Request) {

	seasons, err := c.datasource.GetSeasons()

	if c.HandleError(err, w) {
		return
	}
	c.SendJSON(w, seasons, http.StatusOK)
}
