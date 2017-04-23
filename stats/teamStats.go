package stats

import (
	"github.com/TSAP-Laval/common"
	"github.com/TSAP-Laval/models"
)

// TeamStats représente les statistiques d'une équipe pour
// une saison.
type TeamStats struct {
	ID      uint           `json:"id"`
	Name    string         `json:"name"`
	Players []playerSeason `json:"players"`
}

// GetTeamStats calcule et retourne les statistiques d'une équipe
// pour une saison
func GetTeamStats(teamID uint, seasonID uint, data common.IDatasource) (*TeamStats, error) {

	// On récupère l'équipe sélectionnée.
	t, err := data.GetTeam(teamID)
	if err != nil {
		return nil, err
	}

	//On récupère tous les matchs d'une équipe.
	matches, err := data.GetMatches(teamID, seasonID)

	var nbMatchs = float64(len(*matches))

	if err != nil {
		return nil, err
	}
	// On crée un tableau de la longueur de players
	//dans lequel on fera notre calcul de stats.
	players := make([]playerSeason, len(t.Joueurs))

	metricsList, err := data.GetMetrics(teamID)

	if err != nil {
		return nil, err
	}

	// Les métrics calculée.
	metricSums := make(map[uint]metric)

	// TODO: T'Étais rendu ici

	// On boucle sur tous les joueurs d'une équipe.
	for i, player := range t.Joueurs {
		// On boucle sur tous les matchs
		for _, match := range *matches {

			computedMetrics, err := computeMetrics(&player, &match, metricsList)

			if err != nil {
				return nil, err
			}

			for _, m := range computedMetrics {
				if _, ok := metricSums[m.ID]; ok {
					metricSums[m.ID].Value += m.Value
				}
			}

			// On fait la somme des metrics:
			//Volume de jeu.
			metric1 += m[0].Value
			//Indice d'efficacité.
			metric2 += m[1].Value
			//SCore de performance.
			metric3 += m[2].Value
		}

		var latestMatch *models.Partie
		latestMatch, err = data.GetLatestMatch(teamID)

		if err != nil {
			return nil, err
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
