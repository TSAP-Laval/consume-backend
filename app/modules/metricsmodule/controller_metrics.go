package metricsmodule

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/TSAP-Laval/common"
	"github.com/TSAP-Laval/consume-backend/app/core"
	"github.com/gorilla/mux"
)

// MetricsController gère les requêtes de gestion des métriques
type MetricsController struct {
	core.Controller
	datasource common.IDatasource
	config     *core.ConsumeConfiguration
	logger     *log.Logger
}

// NewMetricsController instancie un MetricsController
func NewMetricsController(datasource common.IDatasource, config *core.ConsumeConfiguration, logger *log.Logger) *MetricsController {
	return &MetricsController{
		Controller: core.Controller{},
		datasource: datasource,
		config:     config,
		logger:     logger,
	}
}

// CreateMetric gère la création d'une métrique
func (c *MetricsController) CreateMetric(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	teamIDRaw := vars["teamID"]

	teamID, err := strconv.Atoi(teamIDRaw)

	if err != nil {
		c.SendJSON(w, core.ErrorMessage{
			Error: fmt.Sprintf("TeamID %s invalid", teamIDRaw),
		}, http.StatusBadRequest)
	}

	var metric MetricsCreationSchema
	if err := c.GetContent(&metric, r); err != nil {
		return
	}

	if metric.Name == "" || metric.Formula == "" {
		c.SendJSON(w, core.ErrorMessage{
			Error: "Invalid Payload",
		}, http.StatusBadRequest)
		return
	}

	err = c.datasource.CreateMetric(metric.Name, metric.Formula, uint(teamID))

	if c.HandleError(err, w) {
		return
	}

	c.SendJSON(w, core.SimpleMessage{
		Body: "ok",
	}, http.StatusCreated)
}
