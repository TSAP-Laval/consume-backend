package stats

import (
	"time"
)

type team struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type metric struct {
	ID        uint    `json:"id"`
	Name      string  `json:"name"`
	Value     float64 `json:"value"`
	Deviation float64 `json:"deviation"`
	LastMatch float64 `json:"last_match"`
	Standard  float64 `json:"standard"`
}

type playerMatch struct {
	ID           uint     `json:"match_id"`
	Date         string   `json:"date"`
	OpposingTeam team     `json:"opposing"`
	Metrics      []metric `json:"metrics"`
}

type season struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type playerSeason struct {
	ID        uint     `json:"id"`
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Metrics   []metric `json:"metrics"`
}

type coach struct {
	ID uint `json:"id"`
}

type typeAction struct {
	Name string `json:"name"`
}

type action struct {
	ID         uint `json:"id"`
	TypeAction typeAction
	IsValid    bool          `json:"is_valid"`
	X1         float64       `json:"x1"`
	Y1         float64       `json:"y1"`
	X2         float64       `json:"x2"`
	Y2         float64       `json:"y2"`
	HomeScore  int           `json:"home_score"`
	AdvScore   int           `json:"adv_score"`
	Time       time.Duration `json:"time"`
}
