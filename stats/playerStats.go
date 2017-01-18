package stats

import (
	"github.com/TSAP-Laval/common"
	"github.com/TSAP-Laval/models"
)

// PlayerStats représente les statistiques d'un joueur pour
// une saison
type PlayerStats struct {
	ID        uint          `json:"player_id"`
	FirstName string        `json:"firstname"`
	LastName  string        `json:"lastname"`
	Team      team          `json:"team"`
	Matches   []playerMatch `json:"matches"`
}

// GetPlayerStats calcule et retourne les statistiques d'un joueur
// pour une saison
func GetPlayerStats(playerID uint, teamID uint, seasonID uint, data *common.Datasource) (*PlayerStats, error) {
	var err error

	// On récupère le joueur
	player, err := data.GetPlayer(playerID)

	if err != nil {
		return nil, err
	}

	// On récupère l'équipe du joueur
	t, err := data.GetTeam(teamID)
	if err != nil {
		return nil, err
	}

	matches, err := data.GetMatches(playerID, teamID, seasonID)

	if err != nil {
		return nil, err
	}

	playerMatches := make([]playerMatch, len(matches))

	for i, match := range matches {
		var advTeam models.Equipe

		if t.ID == match.EquipeMaison.ID {
			advTeam = match.EquipeAdverse
		} else {
			advTeam = match.EquipeMaison
		}

		playerMatches[i] = playerMatch{
			ID:           match.ID,
			Date:         match.Date,
			OpposingTeam: team{ID: advTeam.ID, Name: advTeam.Nom},
			Metrics:      getMetrics(player, &match),
		}
	}

	stats := PlayerStats{
		ID:        playerID,
		FirstName: player.Prenom,
		LastName:  player.Nom,
		Team:      team{ID: teamID, Name: t.Nom},
		Matches:   playerMatches,
	}

	return &stats, err
}
