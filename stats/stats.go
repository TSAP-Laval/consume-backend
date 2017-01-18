package stats

type team struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type metric struct {
	ID        uint    `json:"id"`
	Name      string  `json:"name"`
	Value     float64 `json:"value"`
	Deviation float64 `json:"deviation"`
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
