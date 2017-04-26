package stats

import (
	"math/rand"
	"testing"

	"github.com/TSAP-Laval/models"
)

const ActionCount = 50

var ActionTypes = []models.TypeAction{
	models.TypeAction{Nom: "BC"},
	models.TypeAction{Nom: "BR"},
	models.TypeAction{Nom: "PO"},
	models.TypeAction{Nom: "TB"},
	models.TypeAction{Nom: "PB"},
}

var Metrics = []models.Metrique{
	models.Metrique{
		Nom:      "Volume de jeu",
		Equation: "BC + BR",
		EquipeID: 42,
	},
	models.Metrique{
		Nom:      "Indice d'efficacité",
		Equation: "(BC + PO + TB)/(BC + BR)",
		EquipeID: 42,
	},
	models.Metrique{
		Nom:      "Score de performance",
		Equation: "(BC + BR) + ((BC + PO + TB)/(BC + BR))",
		EquipeID: 42,
	},
}

func TestMetrics(t *testing.T) {

	// Création d'un joueur pour lequel calculer les différentes métriques
	j := &models.Joueur{}
	j.ID = 1337

	// Création d'une cinquantaine d'actions
	mockMatch := &models.Partie{}
	mockMatch.Actions = make([]models.Action, ActionCount)
	for i := 0; i < ActionCount; i++ {
		mockMatch.Actions[i] = models.Action{}

		act := ActionTypes[rand.Intn(len(ActionTypes))]

		var playerID int

		if i%2 == 0 {
			playerID = 1337
		} else {
			playerID = 123
		}

		mockMatch.Actions[i].JoueurID = playerID
		mockMatch.Actions[i].ActionPositive = rand.Intn(10)%2 == 0
		mockMatch.Actions[i].TypeAction = act
	}

	t.Run("computeMetrics raises no error", func(t *testing.T) {
		_, err := computeMetrics(j, mockMatch, &Metrics, &ActionTypes)

		if err != nil {
			t.Errorf("computeMetrics() raised an error: %s", err)
		}
	})

}
