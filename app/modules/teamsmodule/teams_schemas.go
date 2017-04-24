package teamsmodule

import "github.com/TSAP-Laval/consume-backend/stats"

// MapParamsCreationSchema représente le JSON posté vers
// le serveur lors de la modification des paramètres de la map
type MapParamsCreationSchema struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

// MapParamsDisplaySchema est le schéma utilisé pour sérialiser les paramètres
// de la map
type MapParamsDisplaySchema struct {
	ID     uint `json:"id"`
	Width  int  `json:"width"`
	Height int  `json:"height"`
}

// MatchesDisplaySchema représente les infos de base sur un match
type MatchesDisplaySchema struct {
	ID            uint   `json:"id"`
	Lieu          string `json:"lieu"`
	EquipeAdverse string `json:"equipe_adverse"`
	// À ajouter : le pointage?
}

//MatchActions représente l'id du match et tous les joueurs
type MatchActions struct {
	MatchID uint                      `json:"match_id"`
	TeamID  uint                      `json:"team_id"`
	Players []*stats.PlayerMatchStats `json:"players"`
}
