package main

import (
	"fmt"

	"github.com/tsap-laval/consume-backend/consume"
)

func main() {
	fmt.Println("Création des tables...")
	consume.SeedData()
}
