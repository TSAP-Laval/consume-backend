package stats

import "github.com/TSAP-Laval/common"

// PlayerMatchStats représente les actions d'un joueur pour
// un match
type PlayerMatchStats struct {
	ID      uint     `json:"match_id"`
	Date    string   `json:"date"`
	Actions []action `json:"matches"`
}

// GetPlayerActions retourne les actions d'un joueur lors d'un partie
// avec des informations sur celle-ci
func GetPlayerActions(playerID uint, matchID uint, data *common.Datasource) (*PlayerMatchStats, error) {

	// On récupère le joueur
	player, err := data.GetPlayer(playerID)

	if err != nil {
		return nil, err
	}
	if player == nil {
		return nil, nil
	}
	// On récupère le match
	match, err := data.GetMatch(matchID)
	if err != nil {
		return nil, err
	}
	playerActions := make([]action, 0)
	for _, act := range match.Actions {
		if uint(act.JoueurID) == playerID {
			playerAction := action{
				ID:         act.ID,
				TypeAction: typeAction{Name: act.TypeAction.Nom},
				IsValid:    act.ActionPositive,
				X1:         act.X1,
				Y1:         act.Y1,
				X2:         act.X2,
				Y2:         act.Y2,
				HomeScore:  act.PointageMaison,
				AdvScore:   act.PointageAdverse,
				Time:       act.Temps,
			}
			playerActions = append(playerActions, playerAction)
		}
	}

	stats := PlayerMatchStats{
		ID:      match.ID,
		Date:    match.Date,
		Actions: playerActions,
	}

	return &stats, err
}
