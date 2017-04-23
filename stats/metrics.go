package stats

import "github.com/TSAP-Laval/models"
import "github.com/Knetic/govaluate"
import "errors"

// getRuntimeContext récupère le nombre d'actions de chaque type, ce qui permet de substituer
// les types par des nombres au sein des métriques définies par les utilisateurs
func getRuntimeContext(player *models.Joueur, match *models.Partie) map[string]interface{} {

	context := make(map[string]int)

	for _, a := range match.Actions {
		if a.JoueurID == int(player.ID) {
			if _, ok := context[a.TypeAction.Nom]; ok {
				context[a.TypeAction.Nom]++
			} else {
				context[a.TypeAction.Nom] = 0
			}
		}
	}

	// Pas sur de comprendre pourquoi on peut pas cast map[string]int => map[string]interface{}...
	genericContext := make(map[string]interface{}, len(context))

	for k, v := range context {
		genericContext[k] = v
	}

	return genericContext
}

func computeMetrics(player *models.Joueur, match *models.Partie, metrics *[]models.Metrique) ([]metric, error) {

	context := getRuntimeContext(player, match)

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

		// TODO: Remove
		if !ok {
			return nil, errors.New("Bruh")
		}

		computedMetrics[i] = metric{ID: met.ID, Name: met.Nom, Value: fResult}
	}

	return computedMetrics, nil
}
