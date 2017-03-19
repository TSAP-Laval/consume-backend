package metricsmodule

// MetricsCreationSchema représente le JSON posté vers
// le serveur lors de la création d'une métrique
type MetricsCreationSchema struct {
	Name    string `json:"nom"`
	Formula string `json:"equation"`
}
