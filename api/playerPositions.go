package api

import (
	"fmt"
	"net/http"

	"strconv"

	"encoding/json"

	"github.com/gorilla/mux"
)

// PlayerPositionsHandler retourne la liste de positions occupées par le joueur
func (c *ConsumeService) PlayerPositionsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	player := vars["playerID"]

	playerID, err := strconv.Atoi(player)

	if err != nil {
		c.Error(fmt.Sprintf("PlayerID %s invalid.", player))
		w.WriteHeader(http.StatusBadRequest)
		c.ErrorWrite("PlayerID is invalid", w)
		return
	}

	positions, err := c.datasource.GetPositions(playerID)

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

	bytes, err := json.Marshal(positions)

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
