package main

import (
	"log"
	"net/http"

	"github.com/TSAP-Laval/consume-backend/consume/api"
)

func main() {

	router := api.GetRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
