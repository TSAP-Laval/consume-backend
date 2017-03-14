package stats

import (
	"math/rand"
	"testing"

	"github.com/TSAP-Laval/models"
)

const ActionCount = 50

func TestMetrics(t *testing.T) {

	// Création d'un joueur pour lequel calculer les différentes métriques
	j := &models.Joueur{}
	j.ID = 1337

	// Création d'une cinquantaine d'actions
	mockMatch := &models.Partie{}
	mockMatch.Actions = make([]models.Action, ActionCount)
	for i := 0; i < ActionCount; i++ {
		mockMatch.Actions[i] = models.Action{}

		var playerID int

		if i%2 == 0 {
			playerID = 1337
		} else {
			playerID = 123
		}

		mockMatch.Actions[i].JoueurID = playerID
		mockMatch.Actions[i].ActionPositive = rand.Intn(10)%2 == 0
	}

	t.Run("getVJ() result is correct", func(t *testing.T) {

		var bc, br float64

		for _, a := range mockMatch.Actions {
			if a.JoueurID == int(j.ID) {
				switch a.TypeAction.Nom {
				case "BC":
					bc++
				case "BR":
					br++
				}
			}
		}

		expected := bc + br + 1
		actual := getVJ(j, mockMatch)

		if actual != expected {
			t.Errorf("Expected %f, got %f.", expected, actual)
		}
	})

	t.Run("getIE() result is correct", func(t *testing.T) {
		var pos float64

		var vj float64
		vj = 1

		for _, a := range mockMatch.Actions {
			if a.JoueurID == int(j.ID) && a.ActionPositive {
				pos++
			}
		}

		expected := pos / vj
		actual := getIE(j, mockMatch, vj)

		if actual != expected {
			t.Errorf("Expected %f, got %f.", expected, actual)
		}
	})

	t.Run("getSP() result is correct", func(t *testing.T) {

		for i := 0; i < 30; i++ {

			vj := rand.Float64() * 10
			ie := rand.Float64() * 10

			expected := vj + ie
			actual := getSP(vj, ie)

			if actual != expected {
				t.Errorf("Expected %f, got %f.", expected, actual)
			}
		}
	})

	// TODO: Tous les tests de getMetrics() devront être refaits
	// après l'implémentation des métriques dynamiques.

	t.Run("getMetrics() returns three metrics", func(t *testing.T) {

		actual := len(getMetrics(j, mockMatch))

		if actual != 3 {
			t.Errorf("Expected %d metrics, got %d metrics", 3, actual)
		}
	})

	t.Run("getMetrics() first metric is valid", func(t *testing.T) {
		actual := getMetrics(j, mockMatch)[0]

		if actual.ID != 1 || actual.Name != "Volume de Jeu" || actual.Value != getVJ(j, mockMatch) {
			t.Errorf("Metric 1 is incorrect")
		}
	})

	t.Run("getMetrics() second metric is valid", func(t *testing.T) {
		actual := getMetrics(j, mockMatch)[1]

		if actual.ID != 2 || actual.Name != "Indice d'efficacité" || actual.Value != getIE(j, mockMatch, getVJ(j, mockMatch)) {
			t.Errorf("Metric 2 is incorrect")
		}
	})

	t.Run("getMetrics() third metric is valid", func(t *testing.T) {
		actual := getMetrics(j, mockMatch)[2]

		if actual.ID != 3 || actual.Name != "Score de performance" || actual.Value != getSP(getVJ(j, mockMatch), getIE(j, mockMatch, getVJ(j, mockMatch))) {
			t.Errorf("Metric 3 is incorrect")
		}
	})

}
