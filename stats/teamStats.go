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
	// dans lequel on fera notre calcul de stats.
	players := make([]playerSeason, len(t.Joueurs))

	metricsList, err := data.GetMetrics(teamID)
	actionTypes, err := data.GetTypeActions()

	if err != nil {
		return nil, err
	}

	// Les métrics calculée.
	metricSums := make(map[uint]float64)
	metricData := make(map[uint]models.Metrique)

	// On boucle sur tous les joueurs d'une équipe.
	for i, player := range t.Joueurs {

		for _, m := range *metricsList {
			metricSums[m.ID] = 0
			metricData[m.ID] = m
		}

		// On boucle sur tous les matchs
		for _, match := range *matches {

			computedMetrics, err := computeMetrics(&player, &match, metricsList, actionTypes)

			if err != nil {
				return nil, err
			}

			for _, m := range computedMetrics {
				if _, ok := metricSums[m.ID]; ok {
					metricSums[m.ID] += m.Value
				}
			}
		}

		var latestMatch *models.Partie
		latestMatch, err = data.GetLatestMatch(teamID)

		if err != nil {
			return nil, err
		}

		latestMetricsList, err := computeMetrics(&player, latestMatch, metricsList, actionTypes)

		if err != nil {
			return nil, err
		}

		latestMetricsData := make(map[uint]float64)
		for _, latestMetric := range latestMetricsList {
			latestMetricsData[latestMetric.ID] = latestMetric.Value
		}

		displayMetrics := []metric{}

		for k, v := range metricData {
			displayMetrics = append(displayMetrics, metric{ID: v.ID, Name: v.Nom, Value: metricSums[k] / nbMatchs, Deviation: 1, LastMatch: latestMetricsData[k]})
		}

		players[i] = playerSeason{
			ID:        player.ID,
			FirstName: player.Prenom,
			LastName:  player.Nom,
			Metrics:   displayMetrics,
		}
	}

	teamStats := TeamStats{
		ID:      teamID,
		Name:    t.Nom,
		Players: players,
	}

	return &teamStats, err
}
