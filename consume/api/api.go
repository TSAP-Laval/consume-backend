package api

import (
	"io"
	"log"

	"net/http"

	"github.com/gorilla/mux"
)

// ConsumeService represents a single service instance
type ConsumeService struct {
	logger *log.Logger
	Router *mux.Router
}

// New creates a new API service instance
func New(writer io.Writer) *ConsumeService {

	return &ConsumeService{
		logger: log.New(writer, "[consume-api] ", log.Flags()),
		Router: getRouter(),
	}
}

// Start starts the service
func (s *ConsumeService) Start(addr string) error {
	s.logger.Printf("TSAP-Consume started on %s... \n", addr)
	err := http.ListenAndServe(addr, s.Router)
	return err
}
