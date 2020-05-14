package main

import (
	"gitlab.com/Tobbeman/script-runner/internal/config"
	"log"
)

func main(){
	conf, err := config.Setup(config.DefaultPath)
	if err != nil {
		log.Fatal("error opening config: ", err)
	}
	log.Print("Loaded config, token is ", conf.Token)
}


