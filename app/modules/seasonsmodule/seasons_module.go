package seasonsmodule

import (
	"log"
	"net/http"

	"github.com/TSAP-Laval/common"
	"github.com/TSAP-Laval/consume-backend/app/core"
)

// SeasonsModule gère la ressource /season
type SeasonsModule struct {
	routes []core.Route
}

// NewSeasonsModule instancie un nouveau module de gestion des saisons
func NewSeasonsModule(datasource common.IDatasource, config *core.ConsumeConfiguration, logger *log.Logger) *SeasonsModule {

	kc := NewSeasonsController(datasource, config, logger)

	r := []core.Route{
		core.Route{
			Method:  http.MethodGet,
			Path:    "/seasons",
			Handler: kc.GetAllSeasons,
		},
	}

	return &SeasonsModule{
		routes: r,
	}
}

// GetRoutes implémente l'interface core.Module
func (m *SeasonsModule) GetRoutes() []core.Route {
	return m.routes
}
