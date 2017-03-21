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

// GetMapParameters gère la récupération de la taille de la map
func (c *MetricsController) GetMapParameters(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	teamIDRaw := vars["teamID"]

	teamID, err := strconv.Atoi(teamIDRaw)

	if err != nil {
		c.SendJSON(w, core.ErrorMessage{
			Error: fmt.Sprintf("TeamID %s invalid", teamIDRaw),
		}, http.StatusBadRequest)
	}
	parameters, err := c.datasource.GetMapSize(uint(teamID))

	if c.HandleError(err, w) {
		return
	}

	var displayParams = MapParamsDisplaySchema{
		ID:     parameters.ID,
		Width:  parameters.Longeur,
		Height: parameters.Largeur,
	}

	c.SendJSON(w, displayParams, http.StatusOK)
}

// SetMapParameters modifie la préférence de la map pour une équipe
func (c *MetricsController) SetMapParameters(w http.ResponseWriter, r *http.Request) {
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

// GetMetrics gère la récupération des métriques d'une équipe
func (c *MetricsController) GetMetrics(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	teamIDRaw := vars["teamID"]

	teamID, err := strconv.Atoi(teamIDRaw)

	if err != nil {
		c.SendJSON(w, core.ErrorMessage{
			Error: fmt.Sprintf("TeamID %s invalid", teamIDRaw),
		}, http.StatusBadRequest)
	}

	metrics, err := c.datasource.GetMetrics(uint(teamID))

	if c.HandleError(err, w) {
		return
	}

	displayMetrics := make([]MetricsDisplaySchema, len(*metrics))

	for i, m := range *metrics {
		displayMetrics[i] = MetricsDisplaySchema{
			ID:          m.ID,
			Name:        m.Nom,
			Description: m.Description,
			Formula:     m.Equation,
		}
	}

	c.SendJSON(w, displayMetrics, http.StatusOK)
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

	if metric.Name == "" || metric.Formula == "" || metric.Description == "" {
		c.SendJSON(w, core.ErrorMessage{
			Error: "Invalid Payload",
		}, http.StatusBadRequest)
		return
	}

	err = c.datasource.CreateMetric(metric.Name, metric.Formula, metric.Description, uint(teamID))

	if c.HandleError(err, w) {
		return
	}

	c.SendJSON(w, core.SimpleMessage{
		Body: "ok",
	}, http.StatusCreated)
}
