package main

import (
	"log"
	"os"
	"tcmlabs.fr/hexagonal_architecture_golang/internal/user/primary_adapter/cli"
)

// cobra cli that create and get all users

func main() {
	if os.Getenv("TMBD_TOKEN") == "" {
		log.Fatalln("Please set env variable TMBD_TOKEN")
	}
	cmd := cli.InitCli()
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
