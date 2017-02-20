package api

import (
	"encoding/json"
	"net/http"
)

// SeasonsHandler retourne la liste de saisons
func (c *ConsumeService) SeasonsHandler(w http.ResponseWriter, r *http.Request) {

	seasons, err := c.datasource.GetSeasons()

	if err != nil {
		var errMsg string
		c.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)

		if c.config.Debug {
			errMsg = err.Error()
		} else {
			errMsg = "An error occured"
		}

		c.ErrorWrite(errMsg, w)
		return
	}

	bytes, err := json.Marshal(seasons)

	if err != nil {
		var errMsg string
		c.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)

		if c.config.Debug {
			errMsg = err.Error()
		} else {
			errMsg = "An error occured"
		}

		c.ErrorWrite(errMsg, w)
		return
	}

	_, err = w.Write(bytes)

	if err != nil {
		c.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
}
