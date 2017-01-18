package stats

import "github.com/TSAP-Laval/models"

func getVJ(player *models.Joueur, match *models.Partie) float64 {
	var bc, br float64

	for _, a := range match.Actions {
		if a.JoueurID == int(player.ID) {
			switch a.TypeAction.Nom {
			case "BC":
				bc++
			case "BR":
				br++
			}
		}
	}

	return bc + br + 1
}

func getIE(player *models.Joueur, match *models.Partie, vj float64) float64 {
	var pos float64

	for _, a := range match.Actions {
		if a.JoueurID == int(player.ID) && a.ActionPositive {
			pos++
		}
	}

	return pos / vj
}

func getMetrics(player *models.Joueur, match *models.Partie) []metric {
	vj := getVJ(player, match)
	ie := getIE(player, match, vj)

	sp := vj + ie

	return []metric{
		metric{ID: 1, Name: "Volume de Jeu", Value: vj, Deviation: 1},
		metric{ID: 2, Name: "Indice d'efficacité", Value: ie, Deviation: 1},
		metric{ID: 3, Name: "Score de performance", Value: sp, Deviation: 1},
	}
}
