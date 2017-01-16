package main

import "github.com/tsap-laval/consume-backend/consume/api"

import "os"

func main() {
	service := api.New(os.Stdout)
	service.Start(":8080")
}
