package metricsmodule

// MetricsCreationSchema représente le JSON posté vers
// le serveur lors de la création d'une métrique
type MetricsCreationSchema struct {
	Name        string `json:"name"`
	Formula     string `json:"formula"`
	Description string `json:"description"`
}
