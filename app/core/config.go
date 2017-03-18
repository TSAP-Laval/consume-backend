package core

// ConsumeConfiguration représente les paramètres
// exposés par l'application
type ConsumeConfiguration struct {
	DatabaseDriver   string
	ConnectionString string
	SeedDataPath     string
	APIURL           string
	Debug            bool
}
