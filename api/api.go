package api

import (
	"io"
	"log"

	"encoding/json"

	"github.com/TSAP-Laval/common"
	"github.com/braintree/manners"
	"github.com/gorilla/mux"
)

// ConsumeConfiguration représente les paramètres
// exposés par l'application
type ConsumeConfiguration struct {
	DatabaseDriver   string
	ConnectionString string
	SeedDataPath     string
	APIURL           string
	Debug            bool
}

// ConsumeService represents a single service instance
type ConsumeService struct {
	logger     *log.Logger
	datasource *common.Datasource
	config     *ConsumeConfiguration
	server     *manners.GracefulServer
}

// New crée une nouvelle instance du service
func New(writer io.Writer, config *ConsumeConfiguration) *ConsumeService {

	return &ConsumeService{
		logger:     log.New(writer, "[consume-api] ", log.Flags()),
		datasource: common.NewDatasource(config.DatabaseDriver, config.ConnectionString),
		config:     config,
		server:     manners.NewServer(),
	}
}

// Info écrit un message vers le logger du service
func (c *ConsumeService) Info(message string) {
	c.logger.Println(message)
}

// Info écrit un message d'erreur vers le logger du service
func (c *ConsumeService) Error(message string) {
	c.logger.Printf("ERROR - %s\n", message)
}

// ErrorWrite écrit un message d'erreur en format JSON vers le writer
// passé en paramètre
func (c *ConsumeService) ErrorWrite(message string, w io.Writer) error {
	bytes, err := json.Marshal(errorMessage{Error: message})

	if err != nil {
		return err
	}

	_, err = w.Write(bytes)

	return err
}

func (c *ConsumeService) getRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/stats/player/{playerID}/team/{teamID}", c.PlayerStatsHandler)
	r.HandleFunc("/api/seed", c.SeedHandler)

	return r
}

// Start démarre le service
func (c *ConsumeService) Start() {
	go func() {

		c.server.Addr = c.config.APIURL
		c.server.Handler = c.getRouter()
		c.server.ListenAndServe()
		c.Info("Consume shutting down...")
	}()
	c.logger.Printf("TSAP-Consume started on %s... \n", c.config.APIURL)
}

// Stop arrête le service
func (c *ConsumeService) Stop() {
	c.server.Close()
}
