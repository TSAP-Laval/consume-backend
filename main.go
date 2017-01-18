package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/kelseyhightower/envconfig"
	"github.com/tsap-laval/consume-backend/api"
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
	fmt.Print("Press enter to stop server...")
	reader.ReadString('\n')

	service.Stop()

	if err != nil {
		panic(err)
	}
}
