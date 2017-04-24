package teamsmodule

import (
	"log"
	"net/http"

	"github.com/TSAP-Laval/common"
	"github.com/TSAP-Laval/consume-backend/app/core"
)

// TeamsModule gère la ressource /teams
type TeamsModule struct {
	routes []core.Route
}

// NewTeamsModule instancie un nouveau module de gestion de métriques
func NewTeamsModule(datasource common.IDatasource, config *core.ConsumeConfiguration, logger *log.Logger) *TeamsModule {

	kc := NewTeamsController(datasource, config, logger)

	r := []core.Route{
		core.Route{
			Method:  http.MethodPost,
			Path:    "/teams/{teamID}/map",
			Handler: kc.SetMapParameters,
		},
		core.Route{
			Method:  http.MethodGet,
			Path:    "/teams/{teamID}/map",
			Handler: kc.GetMapParameters,
		},
		core.Route{
			Method:  http.MethodGet,
			Path:    "/teams/{teamID}/season/{seasonID}/matches",
			Handler: kc.GetTeamMatches,
		},
		core.Route{
			Method:  http.MethodGet,
			Path:    "/teams/{teamID}/matches/{matchID}/actions",
			Handler: kc.GetTeamActions,
		},
	}

	return &TeamsModule{routes: r}
}

// GetRoutes implémente l'interface core.Module
func (m *TeamsModule) GetRoutes() []core.Route {
	return m.routes
}
