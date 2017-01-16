package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Message Structure de test.
type Message struct {
	Greating string
	From     []string
}

// IndexHandler first handler.
func IndexHandler(w http.ResponseWriter, r *http.Request) {

	res1D := &Message{
		Greating: "Bonjour!",
		From:     []string{"Loic", "Ndjoyi"}}
	res1B, err := json.Marshal(res1D)

	if err != nil {
		fmt.Println(err)
	}

	_, err = w.Write(res1B)
	if err != nil {
		fmt.Println(err)
	}
}
