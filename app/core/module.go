package core

// Module représente une ressource gérée par l'API
type Module interface {
	GetRoutes() []Route
}
