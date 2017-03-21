package seedmodule

import (
	"log"
	"net/http"

	"github.com/TSAP-Laval/common"
	"github.com/TSAP-Laval/consume-backend/app/core"
)

// SeedModule gère la ressource /seed
type SeedModule struct {
	routes []core.Route
}

// NewSeedModule instancie un nouveau module de gestion des saisons
func NewSeedModule(datasource common.IDatasource, config *core.ConsumeConfiguration, logger *log.Logger) *SeedModule {

	kc := NewSeedController(datasource, config, logger)

	r := []core.Route{
		core.Route{
			Method:  http.MethodGet,
			Path:    "/seed",
			Handler: kc.StartSeed,
		},
	}

	return &SeedModule{
		routes: r,
	}
}

// GetRoutes implémente l'interface core.Module
func (m *SeedModule) GetRoutes() []core.Route {
	return m.routes
}
