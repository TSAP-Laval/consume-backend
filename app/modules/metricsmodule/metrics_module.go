package metricsmodule

import (
	"log"
	"net/http"

	"github.com/TSAP-Laval/common"
	"github.com/TSAP-Laval/consume-backend/app/core"
)

// MetricsModule gère la ressource /metrics
type MetricsModule struct {
	routes []core.Route
}

// NewMetricsModule instancie un nouveau module de gestion de métriques
func NewMetricsModule(datasource common.IDatasource, config *core.ConsumeConfiguration, logger *log.Logger) *MetricsModule {

	kc := NewMetricsController(datasource, config, logger)

	r := []core.Route{
		core.Route{
			Method:  http.MethodPost,
			Path:    "/teams/{teamID}/metrics",
			Handler: kc.CreateMetric,
		},
	}

	return &MetricsModule{routes: r}
}

// GetRoutes implémente l'interface core.Module
func (m *MetricsModule) GetRoutes() []core.Route {
	return m.routes
}
