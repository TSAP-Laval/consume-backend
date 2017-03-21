package teamsmodule

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
