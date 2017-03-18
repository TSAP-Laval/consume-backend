package seedmodule

import (
	"fmt"
	"net/http"

	"github.com/TSAP-Laval/common"
	"github.com/TSAP-Laval/consume-backend/app/core"
)

// SeedController répond au call de seeding
type SeedController struct {
	core.Controller
	datasource common.IDatasource
	config     *core.ConsumeConfiguration
}

// NewSeedController instancie un SeedController
func NewSeedController(datasource common.IDatasource, config *core.ConsumeConfiguration) *SeedController {
	return &SeedController{
		Controller: core.Controller{},
		datasource: datasource,
		config:     config,
	}
}

// StartSeed gère l'endpoint de seeding
func (c *SeedController) StartSeed(w http.ResponseWriter, r *http.Request) {

	// TODO: Rendre le seeding disponible seulement en debug mode
	go c.seedData()

	c.SendJSON(w, core.SimpleMessage{Body: "Seeding started."}, http.StatusOK)
}

func (c *SeedController) seedData() {
	// TODO: Enlever les Println
	fmt.Println("Seeding...")
	err := common.SeedData(c.config.DatabaseDriver, c.config.ConnectionString, c.config.SeedDataPath)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Seeding Complete!")
	}
}
