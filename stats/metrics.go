package stats

import (
	"errors"

	"math"

	"github.com/Knetic/govaluate"
	"github.com/TSAP-Laval/common"
	"github.com/TSAP-Laval/models"
)

// getRuntimeContext récupère le nombre d'actions de chaque type, ce qui permet de substituer
// les types par des nombres au sein des métriques définies par les utilisateurs
func getRuntimeContext(player *models.Joueur, match *models.Partie, actionTypes *[]models.TypeAction) map[string]interface{} {

	context := make(map[string]int)

	for _, t := range *actionTypes {
		context[t.Nom] = 0
	}

	for _, a := range match.Actions {
		if a.JoueurID == int(player.ID) {
			context[a.TypeAction.Nom]++
		}
	}

	// Pas sur de comprendre pourquoi on peut pas cast map[string]int => map[string]interface{}...
	genericContext := make(map[string]interface{}, len(context))

	for k, v := range context {
		genericContext[k] = v
	}

	return genericContext
}

func computeMetrics(player *models.Joueur, match *models.Partie, metrics *[]models.Metrique, actionTypes *[]models.TypeAction) ([]metric, error) {

	context := getRuntimeContext(player, match, actionTypes)

	computedMetrics := make([]metric, len(*metrics))

	for i, met := range *metrics {
		// Création d'un bloc évaluable à partir de l'équation définie par l'utilisateur
		expr, err := govaluate.NewEvaluableExpression(met.Equation)

		if err != nil {
			return nil, err
		}

		result, err := expr.Evaluate(context)

		if err != nil {
			return nil, err
		}

		fResult, ok := result.(float64)

		if !ok {
			return nil, errors.New("64-bit float casting error")
		}

		// < *whistling* >
		if math.IsNaN(fResult) || math.IsInf(fResult, 0) {
			fResult = 0
		}
		// </ *whistling* >

		computedMetrics[i] = metric{ID: met.ID, Name: met.Nom, Value: fResult}
	}

	return computedMetrics, nil
}

// ComputeStandard Retourne la norme pour chacun des metriques.
func ComputeStandard(data *common.AllGamesAllPlayerGivenSeason, pActionTypes *[]models.TypeAction, pMetricsList *[]models.Metrique) (map[uint]float64, error) {

	var err error

	standard := make(map[uint]float64)

	for _, m := range *pMetricsList {
		standard[m.ID] = 0
	}

	// On boucle sur tous les joueurs
	for _, player := range data.Players {

		// On boucle sur tous les matchs
		for _, match := range data.Games {

			// On calcul les métriques de tous les matchs.
			computedMetrics, err := computeMetrics(&player, &match, pMetricsList, pActionTypes)

			if err != nil {
				return nil, err
			}

			// On ajoute la valeur obtenue à la liste qui sera retournée.
			for _, m := range computedMetrics {
				if _, ok := standard[m.ID]; ok {
					standard[m.ID] += m.Value
				}
			}
		}
	}

	var nbJoueurs = float64(len(data.Players))

	for k, v := range standard {
		standard[k] = (v / nbJoueurs)
	}

	return standard, err
}
