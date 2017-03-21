package core

// SimpleMessage représente un message INFO
type SimpleMessage struct {
	Body string `json:"body"`
}

// ErrorMessage permet de renvoyer des erreurs
type ErrorMessage struct {
	Error string `json:"error"`
}
