package api

import (
	"encoding/json"

	"net/http"

	"github.com/tsap-laval/common"
)

// SeedHandler gère l'endpoint de seeding
func (c *ConsumeService) SeedHandler(w http.ResponseWriter, r *http.Request) {

	dbtype := c.config.DatabaseDriver
	connstring := c.config.ConnectionString

	go c.seedData(dbtype, connstring)

	bytes, err := json.Marshal(simpleMessage{Body: "Seeding started."})

	if err != nil {
		c.logger.Printf("ERROR - %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		c.ErrorWrite("An error occured", w)
	}

	_, err = w.Write(bytes)

	if err != nil {
		c.logger.Printf("ERROR - %s", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (c *ConsumeService) seedData(dbType string, connString string) {
	c.Info("Seeding...")
	err := common.SeedData(c.config.DatabaseDriver, c.config.ConnectionString, c.config.SeedDataPath)
	if err != nil {
		c.Error(err.Error())
	} else {
		c.Info("Seeding complete.")
	}
}
