package seedmodule

import (
	"log"
	"net/http"

	"github.com/TSAP-Laval/common"
	"github.com/TSAP-Laval/consume-backend/app/core"
)

// SeedController répond au call de seeding
type SeedController struct {
	core.Controller
	datasource common.IDatasource
	config     *core.ConsumeConfiguration
	logger     *log.Logger
}

// NewSeedController instancie un SeedController
func NewSeedController(datasource common.IDatasource, config *core.ConsumeConfiguration, logger *log.Logger) *SeedController {
	return &SeedController{
		Controller: core.Controller{},
		datasource: datasource,
		config:     config,
		logger:     logger,
	}
}

// StartSeed gère l'endpoint de seeding
func (c *SeedController) StartSeed(w http.ResponseWriter, r *http.Request) {
	go c.seedData()

	c.SendJSON(w, core.SimpleMessage{Body: "Seeding started."}, http.StatusOK)
}

func (c *SeedController) seedData() {
	c.logger.Println("Seeding...")
	err := common.SeedData(c.config.DatabaseDriver, c.config.ConnectionString, c.config.SeedDataPath)
	if err != nil {
		c.logger.Printf("Error - %s", err.Error())
	} else {
		c.logger.Println("Seeding Complete!")
	}
}
