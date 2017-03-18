package statsmodule

import (
	"net/http"

	"github.com/TSAP-Laval/common"
	"github.com/TSAP-Laval/consume-backend/app/core"
)

// StatsModule gère la ressource /stats
type StatsModule struct {
	routes []core.Route
}

// NewStatsModule instancie un nouveau module de gestion des statistiques
func NewStatsModule(datasource common.IDatasource, config *core.ConsumeConfiguration) core.Module {

	kc := NewStatsController(datasource, config)

	r := []core.Route{
		core.Route{
			Method:  http.MethodGet,
			Path:    "/stats/player/{playerID}/team/{teamID}",
			Handler: kc.GetPlayerStats,
		},

		core.Route{
			Method:  http.MethodGet,
			Path:    "/stats/match/{matchID}/player/{playerID}",
			Handler: kc.GetPlayerMatchStats,
		},

		core.Route{
			Method:  http.MethodGet,
			Path:    "/stats/team/{teamID}",
			Handler: kc.GetTeamStats,
		},

		core.Route{
			Method:  http.MethodGet,
			Path:    "/stats/player/{playerID}/positions",
			Handler: kc.GetPlayerPositions,
		},
	}

	return &StatsModule{
		routes: r,
	}
}

// GetRoutes implémente l'interface core.Module
func (m *StatsModule) GetRoutes() []core.Route {
	return m.routes
}
