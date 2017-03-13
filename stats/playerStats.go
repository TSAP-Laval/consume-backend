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
func GetPlayerStats(playerID uint, teamID uint, seasonID uint, positionID uint, data common.IDatasource) (*PlayerStats, error) {
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

	matches, err := data.GetMatches(teamID, seasonID)

	if err != nil {
		return nil, err
	}

	// On filtre les match par position et par joueur. Faire ça via GORM aurait été *vraiment* mieux,
	// mais.. Well.. GORM.
	//  ______    ____    _____     _____   _____  __      __  ______     __  __   ______
	// |  ____|  / __ \  |  __ \   / ____| |_   _| \ \    / / |  ____|   |  \/  | |  ____|
	// | |__    | |  | | | |__) | | |  __    | |    \ \  / /  | |__      | \  / | | |__
	// |  __|   | |  | | |  _  /  | | |_ |   | |     \ \/ /   |  __|     | |\/| | |  __|
	// | |      | |__| | | | \ \  | |__| |  _| |_     \  /    | |____    | |  | | | |____
	// |_|       \____/  |_|  \_\  \_____| |_____|     \/     |______|   |_|  |_| |______|

	var filteredMatches []models.Partie

	filteredMatches = []models.Partie{}

	var pos *models.Position
	for _, match := range matches {
		pos, err = data.GetMatchPosition(int(playerID), int(match.ID))
		if err != nil {
			// Le joueur n'était pas dans la partie
			continue
		}

		if positionID != 0 {
			if pos.ID == positionID {
				filteredMatches = append(filteredMatches, match)
			}
		} else {
			filteredMatches = append(filteredMatches, match)
		}
	}

	playerMatches := make([]playerMatch, len(filteredMatches))

	for i, match := range filteredMatches {
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
