package stats

import "github.com/TSAP-Laval/common"

// TeamStats représente les statistiques d'une équipe pour
// une saison.
type TeamStats struct {
	ID      uint           `json:"id"`
	Name    string         `json:"name"`
	CoachID string         `json:"coach_id"`
	Players []playerSeason `json:"players"`
}

// GetTeamStats calcule et retourne les statistiques d'
// pour une saison
func GetTeamStats(coachID uint, teamID uint, seasonID uint, data *common.Datasource) (*TeamStats, error) {
	var err error

	// On récupère le coach.
	coach, err := data.GetCoach(coachID)

	if err != nil {
		return nil, err
	}

	// On récupère l'équipe sélectionnée.
	t, err := data.GetTeam(teamID)
	if err != nil {
		return nil, err
	}

	// On récupère tous les joueurs d'une équipe
	//selon la saison.
	players, err := data.GetPlayers(teamID, seasonID)

	if err != nil {
		return nil, err
	}

	//On récupère tous les matchs d'une équipe.
	matches, err := data.GetMatches(playerID, teamID, seasonID)

	if err != nil {
		return nil, err
	}
	// On crée un tableau de la longueur de players
	//dans lequel on fera notre calcul de stats.
	teamStats := make([]playerSeason, len(players))

	// Les métrics calculée.
	var metric1, metric2, metric3 float64

	// On boucle sur tous les joueurs d'une équipe.
	for i, player := range players {
		// On boucle sur tous les matchs
		for j, match := range matches {

			m := getMetrics(&player, &match)
			// On fait la somme des metrics:
			//Volume de jeu.
			metric1 += m[0].Value
			//Indice d'efficacité.
			metric2 += m[1].Value
			//SCore de performance.
			metric3 += m[2].Value
		}
		metric := []metric{
			metric{Name: "Volume de Jeu", Value: metric1 / len(matches), Deviation: 1},
			metric{Name: "Indice d'efficacité", Value: metric2 / len(matches), Deviation: 1},
			metric{Name: "Score de performance", Value: metric3 / len(matches), Deviation: 1},
		}

		players := playerSeason{
			ID:        playerID,
			FirstName: player.Prenom,
			LastName:  player.Nom,
			Metrics:   metric,
		}
	}
	teamStats := TeamStats{
		ID:      teamID,
		Name:    player.Prenom,
		CoachID: coachID,
		Players: players,
	}

	return &teamStats, err
}
