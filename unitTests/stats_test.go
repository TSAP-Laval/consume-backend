package unitTests

import (
	"os"
	"testing"

    "github.com/TSAP-Laval/common"

    "github.com/TSAP-Laval/consume-backend/stats"
)


func TestStats(t *testing.T) {

    // On récupère la bonne configuration.
	config, err := common.getConfig()

	if err != nil {
		t.Errorf("Error loading configuration: %s", err.Error())
	}

	// On seed une base de données de test
	err = common.SeedData(config.DatabaseDriver, config.ConnectionString, config.SeedDataPath)

	if err != nil {
		t.Errorf("Unexpected exception on seedData function: %s", err.Error())
	}

    //Le datasource qui va servir aux tests.
	data := common.NewDatasource(config.DatabaseDriver, config.ConnectionString)

    // On va chercher la saison en cours.
    currentSeason, err := data.GetCurrentSeason()
    // Test si une saison courante existe.
    if err != nil {
			t.Errorf("Unexpected exception on getCurrentSeason function: %s", err.Error())
		}
    // Test si la saison courante est la bonne.
    if currentSeason.Annees != "2015-2016" {
			t.Errorf("Expected %s, got %s", "2015-2016", currentSeason.Annees)
		}

        // Find the latest match for the team #3.
    match, err := data.GetLatestMatch(3)
    // Test si le dernier match existe.
    if err != nil {
			t.Errorf("Unexpected exception on getLatestMatch function: %s", err.Error())
		}

    // TEAMSTATS.GO
	teamStatsCases := []common.testCase {
		{TestID: 3, IsNil: false, ExpectID: 3},
		{TestID: 99999, IsNil: true, ExpectID: 1},
	}

	for _, c := range teamStatsCases {
		t.Run("GetTeamStats() doesn't fail", func(t *testing.T) {
			_, err := stats.GetTeamStats(c.TestID, currentSeason.ID, data)

			if !c.IsNil && err != nil {
				t.Errorf("Unexpected exception: %s", err.Error())
			}
		})

		t.Run("GetTeamStats() returns correct team stats.", func(t *testing.T) {
			stats, _ := stats.GetTeamStats(c.TestID, currentSeason.ID, data)

			if !c.IsNil && (stats.ID != c.ExpectID) {
				t.Errorf("Expected team %d, got %d", c.ExpectID, team.ID)
			}
		})

		t.Run("GetTeamStats() returns nil when team not found", func(t *testing.T) {
			stats, err := stats.GetTeamStats(c.TestID, currentSeason.ID, data)

			if c.IsNil && ((stats != nil) || err == nil) {
				t.Errorf("Expected team to be Nil, got ID %d instead", team.ID)
			}
		})
	}

    // PLAYERMATCHSTATS.GO
    playerMatchStatsCases := []common.testCase {
		{TestID: 119, IsNil: false, ExpectID: 119},
		{TestID: 99999, IsNil: true, ExpectID: 1},
	}

	for _, c := range playerMatchStatsCases {
		t.Run("GetPlayerActions() doesn't fail", func(t *testing.T) {
			_, err := stats.GetPlayerActions(c.TestID, match.ID, data)

			if !c.IsNil && err != nil {
				t.Errorf("Unexpected exception: %s", err.Error())
			}
		})

		t.Run("GetPlayerActions() returns correct player actions.", func(t *testing.T) {
			actions, _ := stats.GetPlayerActions(c.TestID, match.ID, data)

			if !c.IsNil && (stats.ID != c.ExpectID) {
				t.Errorf("Expected %d, got %d", c.TestID, team.ID)
			}
		})

		t.Run("GetPlayerActions() returns nil when player not found", func(t *testing.T) {
			actions, err := stats.GetPlayerActions(c.TestID, match.ID, data)

			if c.IsNil && ((stactionsats != nil) || err == nil) {
				t.Errorf("Expected player match stats ID to be Nil, got ID %d instead", actions.ID)
			}
		})
	}

    // PLAYERSTATS.GO
    playerStatsCases := []common.testCase {
		{TestID: 119, IsNil: false, ExpectID: 119},
		{TestID: 99999, IsNil: true, ExpectID: 1},
	}

	for _, c := range playerStatsCases {
		t.Run("GetPlayerActions() doesn't fail", func(t *testing.T) {
			_, err := stats.GetPlayerStats(c.TestID, 3, currentSeason.ID, 1, data)

			if !c.IsNil && err != nil {
				t.Errorf("Unexpected exception: %s", err.Error())
			}
		})

		t.Run("GetPlayerActions() returns correct player stats with the position.", func(t *testing.T) {
			pStats, _ := stats.GetPlayerStats(c.TestID, 3, currentSeason.ID, 1, data)

			if !c.IsNil && (pStats.ID != c.ExpectID) {
				t.Errorf("Expected %d, got %d", c.TestID, pStats.ID)
			}
		})

		t.Run("GetPlayerActions() returns nil when player not found", func(t *testing.T) {
			pStats, err := stats.GetPlayerStats(c.TestID, 3, currentSeason.ID, 1, data)

			if c.IsNil && ((pStats != nil) || err == nil) {
				t.Errorf("Expected player stats ID to be Nil, got ID %d instead", pStats.ID)
			}
		})
	}


	// Teardown de la BD de test
	if config.DatabaseDriver == "sqlite3" {
		os.Remove(config.ConnectionString)
	}
}
