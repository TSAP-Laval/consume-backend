package stats

import (
	"errors"
	"testing"
	"time"

	"github.com/TSAP-Laval/models"
)

type MockDatasource struct {
	shouldMatchFail  bool
	shouldPlayerFail bool
}

// Implémentation mock de l'interface datasource... Faudrait rapetisser l'interface parce
// que tbh ça fait long
func (m *MockDatasource) GetCurrentSeason() (*models.Saison, error)   { return nil, nil }
func (m *MockDatasource) GetSeasons() (*[]models.Saison, error)       { return nil, nil }
func (m *MockDatasource) GetTeam(teamID uint) (*models.Equipe, error) { return nil, nil }

func (m *MockDatasource) GetMatches(teamID uint, seasonID uint) (*[]models.Partie, error) {
	return nil, nil
}
func (m *MockDatasource) GetMatchPosition(playerID uint, matchID uint) (*models.Position, error) {
	return nil, nil
}
func (m *MockDatasource) GetMatchActions(teamID uint, matchID uint) (*models.Partie, error) {
	return nil, nil
}
func (m *MockDatasource) GetMatchesInfos(teamID uint) (*[]models.Partie, error) {
	return nil, nil
}
func (m *MockDatasource) GetPositions(playerID uint) (*[]models.Position, error) { return nil, nil }
func (m *MockDatasource) GetLatestMatch(teamID uint) (*models.Partie, error)     { return nil, nil }
func (m *MockDatasource) GetCoach(coachID uint) (*models.Entraineur, error)      { return nil, nil }
func (m *MockDatasource) CreateMetric(name string, formula string, description string, teamID uint) error {
	return nil
}

func (m *MockDatasource) GetMapSize(teamID uint) (*models.MapParameters, error) { return nil, nil }
func (m *MockDatasource) SetMapSize(width int, height int, teamID uint) error {
	return nil
}

func (m *MockDatasource) GetMetrics(teamID uint) (*[]models.Metrique, error) { return nil, nil }

func (m *MockDatasource) UpdateMetric(metricID uint, name string, formula string, description string) error {
	return nil
}
func (m *MockDatasource) DeleteMetric(metricID uint) error { return nil }

func (m *MockDatasource) GetTypeActions() (*[]models.TypeAction, error) { return nil, nil }

// Fonctions de l'interface IDatasource qui sont pertinentes à nos tests
func (m *MockDatasource) GetMatch(matchID uint) (*models.Partie, error) {
	if m.shouldMatchFail {
		return nil, errors.New("WOAH THERE SON")
	}

	nbActions := 10

	p := &models.Partie{Date: "Hello"}
	p.ID = 42

	p.Actions = make([]models.Action, nbActions)
	// Génération d'actions
	for i := 0; i < nbActions; i++ {
		action := models.Action{
			ActionPositive:  true,
			TypeAction:      models.TypeAction{Nom: "MonAction"},
			X1:              1,
			X2:              1,
			Y1:              1,
			Y2:              1,
			PointageAdverse: 1,
			PointageMaison:  1,
			Temps:           time.Duration(1234),
		}

		action.ID = uint(i)
		if i%2 != 0 {
			action.JoueurID = 1337
		} else {
			action.JoueurID = 1
		}

		p.Actions[i] = action
	}

	return p, nil
}
func (m *MockDatasource) GetPlayer(playerID uint) (*models.Joueur, error) {
	if m.shouldPlayerFail {
		return nil, errors.New("WOAH THERE GIRL")
	}
	p := &models.Joueur{}
	p.ID = 1337

	return p, nil
}

func TestPlayerMatchstats(t *testing.T) {

	t.Run("GetPlayerActions() catches errors on match recovery", func(t *testing.T) {
		mockData := MockDatasource{shouldPlayerFail: false, shouldMatchFail: true}

		_, err := GetPlayerActions(0, 0, &mockData)

		if err.Error() != "WOAH THERE SON" {
			t.Errorf("GetPlayerActions() does not catch errors on match recovery")
		}
	})

	t.Run("GetPlayerActions() catches errors on player recovery", func(t *testing.T) {
		mockData := MockDatasource{shouldPlayerFail: true, shouldMatchFail: false}

		_, err := GetPlayerActions(0, 0, &mockData)

		if err.Error() != "WOAH THERE GIRL" {
			t.Errorf("GetPlayerActions() does not catch errors on player recovery")
		}
	})

	t.Run("GetPlayerActins() returns the correct number of actions", func(t *testing.T) {
		mockData := MockDatasource{shouldMatchFail: false, shouldPlayerFail: false}

		act, _ := GetPlayerActions(1337, 0, &mockData)

		if len(act.Actions) != 5 {
			t.Errorf("Expected %d actions, got %d", 5, len(act.Actions))
		}

	})
}
