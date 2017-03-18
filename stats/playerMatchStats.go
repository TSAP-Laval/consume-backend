package stats

import "github.com/TSAP-Laval/common"

// PlayerMatchStats représente les actions d'un joueur pour
// un match
type PlayerMatchStats struct {
	ID      uint     `json:"match_id"`
	Date    string   `json:"date"`
	Actions []action `json:"actions"`
}

// GetPlayerActions retourne les actions d'un joueur lors d'un partie
// avec des informations sur celle-ci
func GetPlayerActions(playerID uint, matchID uint, data common.IDatasource) (*PlayerMatchStats, error) {

	// On récupère le joueur
	_, err := data.GetPlayer(playerID)

	if err != nil {
		return nil, err
	}
	// On récupère le match
	match, err := data.GetMatch(matchID)
	if err != nil {
		return nil, err
	}
	playerActions := []action{}
	for i := 0; i < len(match.Actions); i++ {
		if uint(match.Actions[i].JoueurID) == playerID {
			playerAction := action{
				ID:         match.Actions[i].ID,
				TypeAction: typeAction{Name: match.Actions[i].TypeAction.Nom},
				IsValid:    match.Actions[i].ActionPositive,
				X1:         match.Actions[i].X1,
				Y1:         match.Actions[i].Y1,
				X2:         match.Actions[i].X2,
				Y2:         match.Actions[i].Y2,
				HomeScore:  match.Actions[i].PointageMaison,
				AdvScore:   match.Actions[i].PointageAdverse,
				Time:       match.Actions[i].Temps,
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
