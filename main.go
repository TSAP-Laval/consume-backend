package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/TSAP-Laval/consume-backend/api"
	"github.com/kelseyhightower/envconfig"
)

func main() {

	// On récupère la configuration
	// de l'environnement & on la passe au service
	var c api.ConsumeConfiguration

	err := envconfig.Process("TSAP", &c)

	if err != nil {
		panic(err)
	}

	service := api.New(os.Stdout, &c)
	service.Start()

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Press enter to stop server...")
	reader.ReadString('\n')

	service.Stop()

	if err != nil {
		panic(err)
	}
}
