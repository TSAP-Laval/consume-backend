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

// UpdateMetric gère la modification d'une métrique existante
func (c *MetricsController) UpdateMetric(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	metricIDRaw := vars["teamID"]

	metricID, err := strconv.Atoi(metricIDRaw)

	if err != nil {
		c.SendJSON(w, core.ErrorMessage{
			Error: fmt.Sprintf("MetricID %s invalid", metricIDRaw),
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

	err = c.datasource.UpdateMetric(uint(metricID), metric.Name, metric.Formula, metric.Description)

	if c.HandleError(err, w) {
		return
	}

	c.SendJSON(w, core.SimpleMessage{
		Body: "ok",
	}, http.StatusOK)
}

// DeleteMetric gère la supression d'une métrique existante
func (c *MetricsController) DeleteMetric(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	metricIDRaw := vars["teamID"]

	metricID, err := strconv.Atoi(metricIDRaw)

	if err != nil {
		c.SendJSON(w, core.ErrorMessage{
			Error: fmt.Sprintf("MetricID %s invalid", metricIDRaw),
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

	err = c.datasource.UpdateMetric(uint(metricID), metric.Name, metric.Formula, metric.Description)

	if c.HandleError(err, w) {
		return
	}

	c.SendJSON(w, core.SimpleMessage{
		Body: "ok",
	}, http.StatusOK)
}
