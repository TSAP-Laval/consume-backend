package stats

import "github.com/TSAP-Laval/common"

// TeamStats représente les statistiques d'une équipe pour
// une saison.
type TeamStats struct {
	ID      uint           `json:"id"`
	Name    string         `json:"name"`
	Players []playerSeason `json:"players"`
}

// GetTeamStats calcule et retourne les statistiques d'une équipe
// pour une saison
func GetTeamStats(teamID uint, seasonID uint, data *common.Datasource) (*TeamStats, error) {
	var err error

	// On récupère l'équipe sélectionnée.
	t, err := data.GetTeam(teamID)
	if err != nil {
		return nil, err
	}

	//On récupère tous les matchs d'une équipe.
	matches, err := data.GetMatches(teamID, seasonID)

	var nbMatchs = float64(len(matches))

	if err != nil {
		return nil, err
	}
	// On crée un tableau de la longueur de players
	//dans lequel on fera notre calcul de stats.
	players := make([]playerSeason, len(t.Joueurs))

	// Les métrics calculée.
	var metric1, metric2, metric3 float64

	// On boucle sur tous les joueurs d'une équipe.
	for i, player := range t.Joueurs {
		// On boucle sur tous les matchs
		for _, match := range matches {

			m := getMetrics(&player, &match)
			// On fait la somme des metrics:
			//Volume de jeu.
			metric1 += m[0].Value
			//Indice d'efficacité.
			metric2 += m[1].Value
			//SCore de performance.
			metric3 += m[2].Value
		}

		latestMatch, err := data.GetLatestMatch(teamID)

		if err != nil {
			//Il faudrait afficher une message d'erreur
		}

		latestMetrics := getMetrics(&player, latestMatch)

		metric := []metric{
			metric{ID: 1, Name: "Volume de Jeu", Value: metric1 / nbMatchs, Deviation: 1, LastMatch: latestMetrics[0].Value},
			metric{ID: 2, Name: "Indice d'efficacité", Value: metric2 / nbMatchs, Deviation: 1, LastMatch: latestMetrics[1].Value},
			metric{ID: 3, Name: "Score de performance", Value: metric3 / nbMatchs, Deviation: 1, LastMatch: latestMetrics[2].Value},
		}

		players[i] = playerSeason{
			ID:        player.ID,
			FirstName: player.Prenom,
			LastName:  player.Nom,
			Metrics:   metric,
		}
		// Remise à zéro des metrics calculés
		metric1 = 0
		metric2 = 0
		metric3 = 0
	}

	teamStats := TeamStats{
		ID:      teamID,
		Name:    t.Nom,
		Players: players,
	}

	return &teamStats, err
}
