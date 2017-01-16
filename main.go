package main

import (
	"fmt"

	"log"
	"net/http"

	"github.com/tsap-laval/consume-backend/consume"

	"github.com/TSAP-Laval/consume-backend/consume/api"
)

func main() {

	router := api.GetRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
	fmt.Println("Création des tables...")
	consume.SeedData()
}
