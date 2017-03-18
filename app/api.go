package app

import (
	"fmt"
	"io"
	"log"

	"net/http"

	"github.com/TSAP-Laval/common"
	"github.com/TSAP-Laval/consume-backend/app/core"
	"github.com/TSAP-Laval/consume-backend/app/modules/seasonsmodule"
	"github.com/TSAP-Laval/consume-backend/app/modules/seedmodule"
	"github.com/TSAP-Laval/consume-backend/app/modules/statsmodule"
	"github.com/braintree/manners"
	"github.com/gorilla/mux"
)

// ConsumeService represents a single service instance
type ConsumeService struct {
	logger     *log.Logger
	datasource common.IDatasource
	config     *core.ConsumeConfiguration
	server     *manners.GracefulServer
}

// New crée une nouvelle instance du service
func New(writer io.Writer, config *core.ConsumeConfiguration) *ConsumeService {

	return &ConsumeService{
		logger:     log.New(writer, "[consume-api] ", log.Flags()),
		datasource: common.NewDatasource(config.DatabaseDriver, config.ConnectionString),
		config:     config,
		server:     manners.NewServer(),
	}
}

// Middleware applique les différents middleware
func (c *ConsumeService) Middleware(h http.Handler) http.Handler {
	// Set CORS
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if c.config.Debug {
			// On ouvre l'accès de l'API si ce dernier est en debug
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers",
				"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		}
		h.ServeHTTP(w, r)
	})
}

func (c *ConsumeService) initModules() []core.Module {
	return []core.Module{
		seasonsmodule.NewSeasonsModule(c.datasource, c.config),
		statsmodule.NewStatsModule(c.datasource, c.config),
		seedmodule.NewSeedModule(c.datasource, c.config),
	}
}

func (c *ConsumeService) getRouter() http.Handler {
	r := mux.NewRouter()

	return c.Middleware(r)
}

// Start démarre le service
func (c *ConsumeService) Start() {
	go func() {

		c.server.Addr = c.config.APIURL
		c.server.Handler = c.getRouter()
		c.server.ListenAndServe()
		fmt.Println("Consume shutting down...")
	}()
	c.logger.Printf("TSAP-Consume started on %s... \n", c.config.APIURL)
}

// Stop arrête le service
func (c *ConsumeService) Stop() {
	c.server.Close()
}
